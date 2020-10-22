import { createApp } from 'vue';
import App from './App.vue';

export const HOST: string = process.env.SERVER_HOST || 'localhost';
export const PORT: number = process.env.SERVER_HTTP_PORT || 9001;

export const ENVOY_HOST: string = process.env.ENVOY_HOST
export const ENVOY_PORT: number = process.env.ENVOY_PORT

export const X_AXIS_NUM = 10;

export const COLORS: string[] = [
    '252, 165, 3',
    '3, 128, 252',
    '3, 252, 86',
    '245, 56, 56',
    '182, 36, 255',
    '255, 36, 211',
];


createApp(App).mount('#app');
