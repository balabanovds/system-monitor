<template>
    <div class="home">
        <user-input />
        <chart
            v-for="(c, i) in charts"
            :key="i"
            :name="c.title"
            :width="size.width"
            :height="size.height"
            :metrics="c.metrics"
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
    computed,
    ComputedRef,
} from 'vue';

import { ChartBase, ChartMetric } from '@/components/types';
import UserInput from '@/components/UserInput.vue';
import Chart from '@/components/Chart.vue';
import { GrpcMetricClient } from '@/traits';
import { getBackground, getColor } from '@/utils';
import { COLORS } from '@/consts';

export default defineComponent({
    name: 'Home',
    setup() {
        const size = reactive({
            width: ref(600),
            height: ref(400),
        });

        const client = GrpcMetricClient.getInstance();
        const { list } = client.infoGetter();

        const charts: ComputedRef<Array<ChartBase>> = computed(() => {
            return list.value.map(m => {
                const metrics: Array<ChartMetric> = m.metricTypes.map(
                    (t, i) => {
                        return {
                            type: t,
                            borderColor: getColor(COLORS[i]),
                            backgroundColor: getBackground(COLORS[i]),
                        };
                    },
                );
                return {
                    title: m.title,
                    metrics,
                };
            });
        });

        onMounted(() => {
            client.getInfo();
        });

        onBeforeUnmount(() => {
            client.closeInfoClient();
        });

        return {
            size,
            charts,
        };
    },

    components: {
        UserInput,
        Chart,
    },
});
</script>
