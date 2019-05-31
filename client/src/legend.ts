
interface Key {
	c: string,
	n: string,
	d: string,
}
const legend :Key[] = [{
	c: 'tile-Floor',
	n: 'Floor',
	d: ''
},{
	c: 'tile-Pit',
	n: 'Pit',
	d: 'Engineered for robot disposal'
},{
	c: 'tile-Repair',
	n: 'Repair',
	d: ''
},{
	c: 'tile-Upgrade',
	n: 'Upgrade',
	d: ''
},{
	c: 'tile-Flag',
	n: 'Flag',
	d: ''
},{
	c: 'tile-Conveyor',
	n: 'Conveyor',
	d: 'Moves one space. Also rotates around bends'
},{
	c: 'tile-ExpressConveyor',
	n: 'ExpressConveyor',
	d: 'Moves two spaces. Also rotates around bends'
},{
	c: 'tile-Pusher',
	n: 'Pusher',
	d: 'Pushes robots into adjacent square'
},{
	c: 'tile-Laser',
	n: 'Laser',
	d: 'Fires a laser until hitting a wall or robot'
},{
	c: 'tile-Spawn',
	n: 'Spawn',
	d: ''
},{
	c: 'tile-Gear',
	n: 'Gear',
	d: 'rotates in the indicated direction. Comes in two directional varieties'

},{
	c: 'tile-OffMap',
	n: 'OffMap',
	d: ''
}]
// wall
// robot
export default legend
