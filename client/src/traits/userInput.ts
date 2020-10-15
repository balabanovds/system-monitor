import { ref } from 'vue';

export const userInput = () => {
    const n = ref(0);
    const m = ref(0);

    return {
        n,
        m,
    };
};
