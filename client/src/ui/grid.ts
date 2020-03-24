export class Grid {
	fullWidth: number
	fullHeight: number
	nItemX: number
	nItemY: number
	tileDim: number
	marginX: number
	marginY: number
	width: number
	height: number

	constructor(wPx, hPx, nItemX, nItemY) {
		this.fullWidth = wPx;
		this.fullHeight = hPx;
		this.nItemX = nItemX;
		this.nItemY = nItemY;

		this.tileDim = Math.floor(Math.min(wPx / (nItemX + 2), hPx / (nItemY + 2)));

		this.marginX = Math.round((this.fullWidth - this.tileDim * nItemX) / 2);
		this.marginY = Math.round((this.fullHeight - this.tileDim * nItemY) / 2);

		this.width = wPx;
		this.height = hPx;
	}

	getRow(yi:number) : number
	{
		let h = this.tileDim;
		return this.marginY + yi * h;
	}

	getCol(xi:number) : number
	{
		let w = this.tileDim;
		let h = this.tileDim;
		return this.marginX + xi * w;
	}

	getItemPx(xi:number, yi:number)
		: {x: number, y: number, width: number, height: number}
	{
		let w = this.tileDim;
		let h = this.tileDim;
		return {
			x: this.marginX + xi * w,
			y: this.marginY + yi * h,
			width: w,
			height: h,
		};
	}
}
