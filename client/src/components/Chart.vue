<template>
    <div>
        <h2>{{ name }}</h2>
        <canvas ref="el" :width="width" :height="height"> </canvas>
    </div>
</template>

<script lang="ts">
import {
    computed,
    defineComponent,
    onMounted,
    reactive,
    ref,
    watch,
    toRefs,
    onBeforeUnmount,
} from 'vue';
import Chart from 'chart.js';
import * as dayjs from 'dayjs';

import { GrpcMetricClient } from '../traits';
import { ChartMetric } from './types';
import { X_AXIS_NUM } from '../consts';
import { takeLast } from '../utils';

export default defineComponent({
    name: 'Chart',
    props: {
        name: {
            type: String,
            required: true,
        },
        metrics: {
            type: Object as () => Array<ChartMetric>,
            required: true,
        },
        width: {
            type: Number,
            default: 800,
        },
        height: {
            type: Number,
            default: 400,
        },
    },
    setup(props) {
        if (props.metrics.length === 0) {
            return;
        }

        const client = GrpcMetricClient.getInstance();
        const { data } = client.streamGetters();
        const chart = reactive({
            el: ref<HTMLCanvasElement | null>(null),
            labels: computed(() => {
                const list = data.value
                    .filter(m => m.type === props.metrics[0].type)
                    .map(m => {
                        const d = dayjs.unix(m.seconds);
                        return d.format('hh:mm:ss');
                    });
                return takeLast(list, X_AXIS_NUM);
            }),
            datasets: computed(() => {
                return props.metrics.map(t => {
                    const filtered = data.value.filter(m => m.type === t.type);
                    const label = filtered[0]?.title || 'undefined';

                    return {
                        label: label,
                        borderColor: t.borderColor,
                        backgroundColor: t.backgroundColor,
                        data: takeLast(
                            filtered.map(m => m.value),
                            X_AXIS_NUM,
                        ),
                    };
                });
            }),
        });

        let lineChart: Chart;
        const renderChart = () => {
            lineChart.data.labels = chart.labels;
            lineChart.data.datasets = chart.datasets;
            lineChart.update();
        };

        onMounted(() => {
            if (!chart.el) {
                return;
            }

            lineChart = new Chart(chart.el, {
                type: 'line',
                options: {
                    animation: {
                        duration: 0,
                    },
                    layout: {
                        padding: 30,
                    },
                },
            });
            renderChart();
        });

        onBeforeUnmount(() => lineChart.destroy());

        watch(data.value, () => {
            renderChart();
        });

        return toRefs(chart);
    },
});
</script>

<style></style>
