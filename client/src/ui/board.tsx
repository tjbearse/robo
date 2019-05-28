import React from 'react'
import ReactDOM from 'react-dom'

import { Tile, TileType, Walls } from '../types/board'
import { Robot, Player } from '../types/player'

interface BoardProps {
	board: Tile[][],
	players:Player[],
	uiInfo,
}
export default function Board({board, players, uiInfo} : BoardProps ) {
	const { colors } = uiInfo
	if (board.length < 1) {
		return (<table id="board" />)
	}
	let inPlayPlayers : Player[] = Object.values(players).reduce(( acc:Player[], player:Player ):Player[] => {
		if (player.robot.config) {
			acc.push(player)
		}
		return acc
	}, [])
	let robotCoords: {[x:number]: {[y:number]: Player}} = inPlayPlayers.reduce(( acc, player ) => {
		let robot = player.robot
		let {X, Y} = robot.config.Location
		if (!acc[X]) {
			acc[X] = {}
		}
		acc[X][Y] = player
		return acc
	}, {})
	let width = board.length
	let height = board[0].length
	return (
		<table id="board"><tbody>
			<tr key="-1">
				{
					range(width+2, -1).map((x)=>(
						<OffMapTile key={`${x},-1`}/>
					))
				}
			</tr>
			{
				range(height).map((y) => {
					return (<tr key={y}>
						<OffMapTile key={`-1,${y}`} />
						{
							range(width).map((x) => {
								let tile = board[x][y]
								let robot = null
								let colorNum = 0
								if (robotCoords[x] && robotCoords[x][y]) {
									let player = robotCoords[x][y]
									robot = player.robot
									colorNum = colors.map[player.name]
								}
								return <Tile key={`${x},${y}`} {...{tile,robot,colorNum}} />
							})
						}
						<OffMapTile key={`${width},${y}`} />
					</tr>)
				})
			}
			<tr key={height}>
				{
					range(width+2, -1).map((x)=>(
						<OffMapTile key={`${x},${height}`}/>
					))
				}
			</tr>
		</tbody></table>
	)
}

function OffMapTile({}) {
	return (
		<td className="tile tile-OffMap"></td>
	)
}

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

// https://stackoverflow.com/questions/3895478/does-javascript-have-a-method-like-range-to-generate-a-range-within-the-supp
function range(size, startAt = 0) {
    return [...Array(size).keys()].map(i => i + startAt);
}
