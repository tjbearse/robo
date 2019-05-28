import React from 'react'
import ReactDOM from 'react-dom'

export default function Error({error, clear}) {
	return (
		<div id="error" onClick={()=>clear()}>{error}</div>
	)
}
