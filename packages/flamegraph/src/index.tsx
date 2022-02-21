/* eslint-disable react/react-in-jsx-scope */
/* eslint-disable react/jsx-props-no-spreading */
import React from 'react';
import FlameGraphRenderer from './FlameGraph/FlameGraphRenderer';
import Flamegraph from './FlameGraph/FlameGraphComponent/Flamegraph';
import ExportData from '../../../webapp/javascript/components/ExportData';
import { decodeFlamebearer } from '../../../webapp/javascript/models/flamebearer';

const overrideProps = {
  showPyroscopeLogo: true,
};

// TODO: type props
const FlamegraphRenderer = (props: any) => {
  return <FlameGraphRenderer {...props} {...overrideProps} />;
};

export { FlamegraphRenderer, Flamegraph, ExportData, decodeFlamebearer };
