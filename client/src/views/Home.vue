<template>
    <div class="home">
        <user-input />
        <chart
            name="Load Average"
            :width="size.width"
            :height="size.height"
            :chartTypes="chartTypes"
        />
    </div>
</template>

<script lang="ts">
import {
    defineComponent,
    reactive,
    ref,
    onMounted,
    onBeforeUnmount,
} from 'vue';
import { ChartType } from '../components/types';
import UserInput from '../components/UserInput.vue';
import Chart from '../components/Chart.vue';
import { MetricType } from '../proto/metric_service_pb';
import { GrpcMetricClient } from '../traits';
import { getBackground, getColor } from '../utils';

export default defineComponent({
    name: 'Home',
    setup() {
        const size = reactive({
            width: ref(600),
            height: ref(400),
        });

        const client = GrpcMetricClient.getInstance();
        const chartTypes = ref<Array<ChartType>>([
            {
                type: MetricType.LOADAVERAGE1MIN,
                borderColor: getColor('20,20,20'),
                backgroundColor: getBackground('20,20,20'),
            },
            {
                type: MetricType.LOADAVERAGE5MIN,
                borderColor: getColor('20,20,20'),
                backgroundColor: getBackground('20,20,20'),
            },
            {
                type: MetricType.LOADAVERAGE15MIN,
                borderColor: getColor('20,20,20'),
                backgroundColor: getBackground('20,20,20'),
            },
        ]);

        onMounted(() => {
            client.getInfo();
        });

        onBeforeUnmount(() => {
            client.closeInfoClient();
        });

        return {
            size,
            chartTypes,
        };
    },

    components: {
        UserInput,
        Chart,
    },
});
</script>
