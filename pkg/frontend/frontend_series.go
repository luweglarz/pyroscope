package frontend

import (
	"context"

	"connectrpc.com/connect"
	"github.com/grafana/dskit/tenant"
	"github.com/opentracing/opentracing-go"
	"github.com/prometheus/common/model"

	querierv1 "github.com/grafana/pyroscope/api/gen/proto/go/querier/v1"
	"github.com/grafana/pyroscope/api/gen/proto/go/querier/v1/querierv1connect"
	querybackendv1 "github.com/grafana/pyroscope/api/gen/proto/go/querybackend/v1"
	phlaremodel "github.com/grafana/pyroscope/pkg/model"
	"github.com/grafana/pyroscope/pkg/querybackend"
	"github.com/grafana/pyroscope/pkg/querybackend/queryplan"
	"github.com/grafana/pyroscope/pkg/util/connectgrpc"
	"github.com/grafana/pyroscope/pkg/validation"
)

func (f *Frontend) Series(ctx context.Context, c *connect.Request[querierv1.SeriesRequest]) (*connect.Response[querierv1.SeriesResponse], error) {
	opentracing.SpanFromContext(ctx).
		SetTag("start", model.Time(c.Msg.Start).Time().String()).
		SetTag("end", model.Time(c.Msg.End).Time().String()).
		SetTag("matchers", c.Msg.Matchers).
		SetTag("label_names", c.Msg.LabelNames)

	ctx = connectgrpc.WithProcedure(ctx, querierv1connect.QuerierServiceSeriesProcedure)
	tenantIDs, err := tenant.TenantIDs(ctx)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	interval, ok := phlaremodel.GetTimeRange(c.Msg)
	if ok {
		validated, err := validation.ValidateRangeRequest(f.limits, tenantIDs, interval, model.Now())
		if err != nil {
			return nil, connect.NewError(connect.CodeInvalidArgument, err)
		}
		if validated.IsEmpty {
			return connect.NewResponse(&querierv1.SeriesResponse{}), nil
		}
		c.Msg.Start = int64(validated.Start)
		c.Msg.End = int64(validated.End)
	}

	labelSelector, err := buildLabelSelectorFromMatchers(c.Msg.Matchers)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	rep, err := f.invoke(ctx, c.Msg.Start, c.Msg.End, tenantIDs, labelSelector, &querybackendv1.Query{
		QueryType: querybackendv1.QueryType_QUERY_SERIES_LABELS,
		SeriesLabels: &querybackendv1.SeriesLabelsQuery{
			LabelNames: c.Msg.LabelNames,
		},
	})
	if err != nil {
		return nil, err
	}
	if rep == nil {
		return connect.NewResponse(&querierv1.SeriesResponse{}), nil
	}
	return connect.NewResponse(&querierv1.SeriesResponse{
		LabelsSet: rep.SeriesLabels.SeriesLabels,
	}), nil
}

func (f *Frontend) invoke(
	ctx context.Context,
	startTime, endTime int64,
	tenants []string,
	labelSelector string,
	q *querybackendv1.Query,
) (*querybackendv1.Report, error) {
	blocks, err := f.listMetadata(ctx, tenants, startTime, endTime, labelSelector)
	if err != nil {
		return nil, err
	}
	// TODO: Params.
	p := queryplan.Build(blocks, 20, 50)
	resp, err := f.querybackendclient.Invoke(ctx, &querybackendv1.InvokeRequest{
		Tenant:        tenants,
		StartTime:     startTime,
		EndTime:       endTime,
		LabelSelector: labelSelector,
		Options:       &querybackendv1.InvokeOptions{},
		QueryPlan:     p.Proto(),
		Query:         []*querybackendv1.Query{q},
	})
	if err != nil {
		return nil, err
	}
	return findReport(querybackend.QueryReportType(q.QueryType), resp.Reports), nil
}
