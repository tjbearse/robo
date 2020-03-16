import robot0 from '../img/intermediates/robot0.png';
import robot1 from '../img/intermediates/robot1.png'
import robot2 from '../img/intermediates/robot2.png'
import robot3 from '../img/intermediates/robot3.png'
import robot4 from '../img/intermediates/robot4.png'
import robot5 from '../img/intermediates/robot5.png'
import robot6 from '../img/intermediates/robot6.png'
import robot7 from '../img/intermediates/robot7.png'

import floor from '../img/intermediates/floor.png'
import pit from '../img/intermediates/pit.png'
import repair from '../img/intermediates/repair.png'
import upgrade from '../img/intermediates/upgrade.png'
import flag from '../img/intermediates/flag.png'
import conveyor from '../img/intermediates/conveyor.png'
import express from '../img/intermediates/express.png'
import pusher from '../img/intermediates/pusher.png'
import laser from '../img/intermediates/laser.png'
import spawn from '../img/intermediates/spawn.png'
import gear from '../img/intermediates/gear.png'

import {TileType} from '../types/board'

export const floorImages = {
	[TileType.Floor]: floor,
	[TileType.Pit]: pit,
	[TileType.Repair]: repair,
	[TileType.Upgrade]: upgrade,
	[TileType.Flag]: flag,
	[TileType.Spawn]: spawn,
	[TileType.Conveyor]: conveyor,
	[TileType.ExpressConveyor]: express,
	[TileType.Pusher]: pusher,
	[TileType.Gear]: gear,
	[TileType.Laser]: laser,
};

export const robotImages = [
	robot0,
	robot1,
	robot2,
	robot3,
	robot4,
	robot5,
	robot6,
	robot7,
];
