

// https://stackoverflow.com/questions/3895478/does-javascript-have-a-method-like-range-to-generate-a-range-within-the-supp
export default function range(size:number, startAt = 0) : number[] {
    return [...Array(size).keys()].map(i => i + startAt);
}
