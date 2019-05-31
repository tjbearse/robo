import React from 'react'
import ReactDOM from 'react-dom'

import { Provider } from 'react-redux'
import store from '../store'
import Layout from './layout'
import ReactModal from 'react-modal'

const rootElement = document.getElementById('gameArea')
ReactModal.setAppElement('#gameArea')
ReactDOM.render(
  <Provider store={store}>
    <Layout />
  </Provider>,
  rootElement
)
