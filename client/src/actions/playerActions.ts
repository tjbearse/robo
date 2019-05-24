import store from '../store'

function ClearError() {
	store.dispatch({
		type: 'ClearError',
	})
}
ClearError.type = 'ClearError'

export {ClearError}
