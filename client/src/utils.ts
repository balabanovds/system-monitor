function takeLast<T>(arr: Array<T>, num: number): Array<T> {
    if (arr.length < num) {
        return arr.slice();
    }
    return arr.slice(arr.length - num, arr.length);
}

export { takeLast };
