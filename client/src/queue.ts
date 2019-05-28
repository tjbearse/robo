import store from './store'
import notify from './actions/notify'

interface action {
	type: string
	payload: any
}

enum Types {
	Robot  = 0,
	Other
}

const timeout = 1000

class Queue {
	messages : action[]
	processInterval : number
	clearInterval : number
	types : {[t : number]: boolean}
	constructor() {
		this.messages = []
		this.processInterval = null
		this.clearInterval = null
		this.types = {}
	}

	push(a : action) {
		this.messages.push(a)
		this.process()
	}

	scheduleProcess() {
		if (!this.processInterval) {
			this.processInterval = setTimeout(() => {
				this.types = {}
				this.processInterval = null
				this.process()
			}, timeout)
		}
	}

	scheduleClear() {
		if (!this.clearInterval) {
			this.clearInterval = setTimeout(() => {
				this.types = {}
				this.clearInterval = null
			}, timeout)
		}
	}

	process() {
		while (this.messages.length > 0) {
			let m = this.messages[0]
			let t = this.getType(m.type)
			if (t != Types.Other) {
				if (this.types[t]) {
					break;
				}
				this.types[t] = true;
			}
			store.dispatch(this.messages.shift())
		}
		if (this.messages.length > 0) {
			this.scheduleProcess()
		} else {
			this.scheduleClear()
		}
	}

	getType(t : string) {
		switch (t) {
		case notify.RobotMoved:
		case notify.RobotFell:
			return Types.Robot
		default:
			return Types.Other
		}
	}
}
let queue = new Queue()
export default queue
