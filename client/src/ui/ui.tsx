import {Walls, TileType, Tile} from "../types/board"
import {ClearError} from '../actions/playerActions'
import {Player} from "../types/player"
import { Card, Command, commandToText } from '../types/card'
// end old imports to break up

import React from 'react'
import ReactDOM from 'react-dom'

import { Provider } from 'react-redux'
import store from '../store'
import Layout from './layout'

const rootElement = document.getElementById('gameArea')
ReactDOM.render(
  <Provider store={store}>
    <Layout />
  </Provider>,
  rootElement
)
