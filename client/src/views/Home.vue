<template>
  <div class="home">
    <div class="error" v-show="disabled">{{ error }}</div>
    <user-input/>
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
import {computed, ComputedRef, defineComponent, onBeforeUnmount, onMounted, reactive, ref,} from 'vue';

import {ChartBase, ChartMetric} from '@/components/types';
import UserInput from '@/components/UserInput.vue';
import Chart from '@/components/Chart.vue';
import {GrpcMetricClient} from '@/traits';
import {getBackground, getColor} from '@/utils';
import {COLORS} from '@/consts';

export default defineComponent({
  name: 'Home',
  setup() {
    const size = reactive({
      width: ref(600),
      height: ref(400),
    });

    const client = GrpcMetricClient.getInstance();
    const {list, error} = client.infoGetter();

    const disabled = computed(()=> error.value !== '')

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
      error,
      disabled
    };
  },

  components: {
    UserInput,
    Chart,
  },
});
</script>

<style>
.error {
  color: red;
  font-weight: bold;
}
</style>
