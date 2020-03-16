import React from 'react'
import ReactDOM from 'react-dom'
import { Stage, Layer, Group, Rect, Text} from 'react-konva';
import useImage from 'use-image';

import { Tile, Directions, TileType, Walls } from '../types/board'
import { Robot, Player } from '../types/player'
import range from './range'
import {UrlImage} from './image'

import {floorImages, robotImages} from './images'

// TODO split wall out from the tiles
interface BoardProps {
	board: Tile[][],
	players: { [name: string] : Player },
	uiInfo,
}

let boardColors = {
	offGrid: '#222',
};

enum DirRotation {
	North = 0,
	East = 1,
	South = 2,
	West = 3,
}

class Grid {
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

export default function Board({board, players, uiInfo} : BoardProps ) {
	if (board.length < 1) return null;
	const grid = new Grid(520, 640, board.length, board[0].length);
	const { colors } = uiInfo
	// TODO walls
	// TODO separate draw for movable items like flags, spawn markers
	// TODO laser fire
    return (
		// @ts-ignore
		<Stage width={600} height={640}>
			<Layer>
				<Map {...{grid, board}} />
				<GridLines grid={grid} />
			</Layer>
			<Layer>
				{
					Object.values(players)
						.filter(p => p.robot.config)
						.map(p => {
							let robot = p.robot;
							let colorNum = colors.map[p.name];
							let key = p.name;
							return (<Robot {...{grid, robot, colorNum, key}} />);
						})
				}
			</Layer>
		</Stage>
	);
}


function GridLines({grid} : {grid : Grid}) {
	return null;
}

function Map({grid, board} : {grid : Grid, board : Tile[][]}) {
	const x = 0;
	const y = 0;
	const width = grid.fullWidth;
	const height = grid.fullHeight;
	return (
		<Group>
			<Rect {...{
					x, y,
					width, height,
					fill: boardColors.offGrid,
			}} />
			{
				// tiles
				board .map((row, xi) => row.map((tile, yi) => {
					const key = 'tile-' + board.length * yi + xi;
					const {x, y, width, height} = grid.getItemPx(xi, yi);
					return (
						<Tile {...{key, x, y, width, height, tile}} />
					);
				}))
			}
		</ Group>
	);
}

// TODO split out tiles into separate functions? Would be nice for animations
function Tile({x,y,width,height,tile}) {
	const image = floorImages[tile.type];
	let fill = tile.type;
	fill = '#' + fill + fill + fill;
	let rotation;
	let offset = { x: width/2, y: height/2 };
	if (tile.dir == null || tile.dir == undefined) {
		// TODO indeterminent animation
		rotation = 0;
	} else {
		rotation = 90 * +DirRotation[tile.dir];
	}
	let scaleX = 1;
	if (tile.type === TileType.Gear && tile.dir !== Directions.North) {
		scaleX = -1;
	}
	return (
		<Group>
			<Rect {...{x,y,width,height, fill}} />
			<UrlImage {...{
				image,
				x: x + offset.x,
				y: y + offset.y,
				width, height,
				rotation,
				offset,
				scaleX,
			}} />
		</Group>
	)
}

function Robot({grid, robot, colorNum}
	: {grid: Grid, robot:Robot, colorNum:number}
) {
	const image = robotImages[colorNum];
	if (!image) return null;

	const {x, y, width, height} = grid.getItemPx(robot.config.Location.X, robot.config.Location.Y);
	const rotation = 0; // TODO
	return (
		<UrlImage {...{
				image,
				x, y,
				width, height,
				rotation,
		}} />
	);
}

/*
function Tile({tile, robot, colorNum}: {tile:Tile, robot:Robot, colorNum:number}) {
	let className = 'tile tile-' + TileType[tile.type] +
					' dir-' + tile.dir
	let text = null
	if (tile.type == TileType.Flag) {
		text = tile.num + 1
	}
	return (
		<td className={className}>{text}
			{ robot && <Robot {...{colorNum, robot}} /> }
			<Wall walls={tile.walls}/>
		</td>
	)
}

function Robot({robot, colorNum}: {robot: Robot, colorNum:number}) {
	let className = 'robot dir-'+(robot.config.Heading || 'indeterminent') +
		' robot-' + colorNum
	return (
		<div className={className}></div>
	)
}

function Wall({walls}: {walls:Walls}) {
	let wallClass = [Walls.North, Walls.East, Walls.South, Walls.West]
		.reduce((acc, w) => {
			if (walls & w) {
				acc += 'wall-' + Walls[w] + ' '
			}
			return acc
		}, '')
	if (!wallClass) {
		return null
	}
	return (<div className={'wall ' + wallClass}></div>)
}
*/
