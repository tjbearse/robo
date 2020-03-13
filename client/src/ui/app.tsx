import React from 'react'
import ReactDOM from 'react-dom'

import { Provider } from 'react-redux'
import store from '../store'

import Layout from './layout'
class Temp extends React.Component {
}

const rootElement = document.getElementById('root')
ReactDOM.render(
	<Provider store={store}>
		<Layout />
	</Provider>,
	rootElement
)

