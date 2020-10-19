function takeLast<T>(arr: Array<T>, num: number): Array<T> {
    if (arr.length < num) {
        return arr.slice();
    }
    return arr.slice(arr.length - num, arr.length);
}

const getColor = (rgbColor: string): string => {
    return `rgba(${rgbColor}, .6)`;
};
const getBackground = (rgbColor: string): string => {
    return `rgba(${rgbColor}, .1)`;
};

export { takeLast, getColor, getBackground };
