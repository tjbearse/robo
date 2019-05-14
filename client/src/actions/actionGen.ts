function createAction<T>(name: string) {
    const fn = function (t: T) {
        return {
            type: name,
            payload: t,
        }
    }
    fn.type = name;
    return fn
}

export default createAction
