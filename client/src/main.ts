import { createApp } from 'vue';
import App from './App.vue';

export const HOST: string = process.env.API_HOST || 'localhost';
export const PORT: number = process.env.API_PORT || 9001;

createApp(App).mount('#app');
