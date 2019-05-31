import React from 'react'
import ReactDOM from 'react-dom'
import ReactModal from 'react-modal'

import legend from '../legend.ts'

export function Help() {
	return (
		<div id="help"><HelpText/></div>
	)
}

function HelpText() {
	return (
		<div>
			<h2>Robot Races</h2>
			<p>
				Be the first to reach the checkpoints in order by sending your robot batches of instructions. Will your carefully laid plans be sucessful or thwarted by obstacles and competing robots? As your robot becomes damaged programming it becomes harder and harder. Watch out!
			</p>
			<p>This help text can be brought up at any time by pressing <span className="help-btn"></span>.</p>
			<GamePlayPhases/>
			<h3>Legend</h3>
			<table>
				<tbody>
					<tr><th>Tile</th><th>Name</th><th>Description</th></tr>
					{
						legend.map(({c,n,d}, i) => (
							<tr key={i}>
								<td className={`tile ${c}`}></td>
								<td>{n}</td>
								<td>{d}</td>
							</tr>
						))
					}
				</tbody>
			</table>
		</div>
	)
}

function GamePlayPhases () {
	return (
		<div>
			<h3>Game Play Phases</h3>
			<div>
				<ol>
					<li><b>Spawn Robots:</b> robots who are destroyed or fall off the map will be replaced by a new copy at their spawn point. Robots start with spawns in predetermined locations. Reaching checkpoints and repair stations will allow your robot to spawn there instead.</li>
					<li><b>Download Instructions:</b> players choose a set of instructions to send to their robot. Commiting a partial set of instructions will force the robot select randomly from your hand!</li>
					<li><b>Run Instructions:</b> robot instructions are run round robin in the following order:
						<ol>
							<li>All Robots run a single instruction. Order between robots is decided by the priority number of their instructions.</li>
							<li>Environmental Hazards activate:
								<ol>
									<li>Conveyor belts move</li>
									<li>Pushers push</li>
									<li>Gears rotate</li>
									<li>Lasers fire: <u>robots are also equipped with front facing lasers!</u></li>
								</ol>
							</li>
							<li>Checkpoints are counted</li>
						</ol>
					</li>
					<li><b>Repairs and upgrades activate:</b> ending the round on a repair space removes 1 damage, upgrades remove 2.</li>
				</ol>
			</div>
		</div>
	)
}

interface HelpHintState {
	expanded: boolean,
}
export class HelpHint extends React.Component {
	state: HelpHintState
	constructor(props) {
		super(props)
		this.state = {
			expanded: false,
		}
	}
	render() {
		const toggle = ()=>{ this.setState({expanded: !this.state.expanded}) }
		return (
			<span onClick={toggle} className="help-btn">
				<ReactModal 
					isOpen={this.state.expanded}
					contentLabel="Help"
				>
					<div className="help-modal" id="help">
						<div className="close" onClick={ toggle }>X</div>
						<HelpText/>
					</div>
				</ReactModal>
			</span>
		)
	}
}
