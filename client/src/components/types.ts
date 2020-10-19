type ChartMetric = {
    type: number;
    borderColor: string;
    backgroundColor: string;
};

type ChartBase = {
    title: string;
    metrics: Array<ChartMetric>;
};

export { ChartBase, ChartMetric };
