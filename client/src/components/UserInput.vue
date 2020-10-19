<template>
    <div>
        <input
            type="text"
            v-model="n"
            :disabled="active"
            placeholder="tick period"
        />
        <input
            type="text"
            v-model="m"
            :disabled="active"
            placeholder="measurement period"
        />
        <button @click="onConnect" v-show="!active">run</button>
        <button @click="onCancel" v-show="active">stop</button>
    </div>
</template>

<script lang="ts">
import { GrpcMetricClient } from '../traits';
import { defineComponent } from 'vue';

export default defineComponent({
    name: 'UserInput',
    setup() {
        const client = GrpcMetricClient.getInstance();
        const { n, m, active } = client.streamGetters();

        function onConnect() {
            client.startStream();
        }

        function onCancel() {
            client.stopStream();
        }

        return {
            n,
            m,
            active,
            onConnect,
            onCancel,
        };
    },
});
</script>

<style></style>
