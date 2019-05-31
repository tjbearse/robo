import React from 'react'

export default class HandleChange extends React.Component {
	constructor(props) {
		super(props)
		this.handleChange = this.handleChange.bind(this)
	}
	handleChange(event) {
		const target = event.target.name
		const value = event.target.value
		this.setState({[target]: value})
	}
}
