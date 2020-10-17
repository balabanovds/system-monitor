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
import { defineComponent, reactive, ref } from 'vue';
import { ChartType } from '../components/types';
import UserInput from '../components/UserInput.vue';
import Chart from '../components/Chart.vue';
import { MetricType } from '../proto/metric_service_pb';

const COLOR_1 = 'rgba(252, 165, 3, .6)';
const BG_COLOR_1 = 'rgba(252, 165, 3, .1)';
const COLOR_2 = 'rgba(3, 128, 252, .6)';
const BG_COLOR_2 = 'rgba(3, 128, 252, .1)';
const COLOR_3 = 'rgba(3, 252, 86, .6)';
const BG_COLOR_3 = 'rgba(3, 252, 86, .1)';

export default defineComponent({
    name: 'Home',
    setup() {
        const size = reactive({
            width: ref(600),
            height: ref(400),
        });
        const chartTypes = ref<Array<ChartType>>([
            {
                type: MetricType.LOADAVERAGE1MIN,
                borderColor: COLOR_1,
                backgroundColor: BG_COLOR_1,
            },
            {
                type: MetricType.LOADAVERAGE5MIN,
                borderColor: COLOR_2,
                backgroundColor: BG_COLOR_2,
            },
            {
                type: MetricType.LOADAVERAGE15MIN,
                borderColor: COLOR_3,
                backgroundColor: BG_COLOR_3,
            },
        ]);

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
