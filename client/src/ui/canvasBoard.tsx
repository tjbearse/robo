import React from 'react'
import ReactDOM from 'react-dom'
import { Stage, Layer, Group, Rect, Text} from 'react-konva';
import useImage from 'use-image';

import { Tile, Board, Directions, TileType, Walls } from '../types/board'
import { Robot, Player } from '../types/player'
import range from './range'
import {UrlImage} from './image'
import {Grid} from './grid'

import {floorImages, robotImages} from './images'

// TODO split wall out from the tiles
interface BoardProps {
	board: Board,
	players: { [name: string] : Player },
	uiInfo,
}

let boardColors = {
	offGrid: '#222',
	wall: '#d4aa00',
};

enum DirRotation {
	North = 0,
	East = 1,
	South = 2,
	West = 3,
}

export default function Board({board, players, uiInfo} : BoardProps ) {
	if (!board.tiles || board.tiles.length < 1) return null;
	const grid = new Grid(520, 640, board.tiles.length, board.tiles[0].length);
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

function Map({grid, board} : {grid : Grid, board : Board}) {
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
				board.tiles.map((row, xi) => row.map((tile, yi) => {
					const key = 'tile-' + board.tiles.length * yi + xi;
					const {x, y, width, height} = grid.getItemPx(xi, yi);
					return (
						<Tile {...{key, x, y, width, height, tile}} />
					);
				}))
			}
			{
				board.nWalls.map((row, xi) => row
					.map((wall, yi) => {
						if (!wall) return null;
						const key = 'nwall-' + board.nWalls.length * yi + xi;
						const {x, y, width, height} = grid.getItemPx(xi, yi);
						return (
							<NWall {...{key, x, y, width, height}} />
						);
					})
					.filter(wall => wall)
				)
			}
			{
				board.wWalls.map((row, xi) => row
					.map((wall, yi) => {
						if (!wall) return null;
						const key = 'wwall-' + board.wWalls.length * yi + xi;
						const {x, y, width, height} = grid.getItemPx(xi, yi);
						return (
							<WWall {...{key, x, y, width, height}} />
						);
					})
					.filter(wall => wall)
				)
			}
		</ Group>
	);
}

function NWall({x,y, width, height}) {
	let fill = boardColors.wall;
	height /= 8;
	y -= height / 2;
	return <Rect {...{x,y,width,height, fill}} />
}

function WWall({x,y, width, height}) {
	let fill = boardColors.wall;
	width /= 8;
	x -= width / 2;
	return <Rect {...{x,y,width,height, fill}} />
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
