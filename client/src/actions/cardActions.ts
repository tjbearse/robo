import createAction from './actionGen'
import { Card } from 'types/card'

/*
interface CardToBoardActionFormat {
	BoardSlot: number;
	HandOffset: number;
	Card: Card;
}
export const CardToBoardAction = createAction<CardToBoardAction>('CardToBoardAction')

interface CardToHandActionFormat {
	BoardSlot: number;
	HandOffset: number;
	Card: Card;
}
export const CardToHandAction = createAction<CardToHandActionFormat>('CardToHandAction')

interface CardToBoardBlindActionFormat {
	Player: string;
	BoardSlot: number;
}
export const CardToBoardBlindAction = createAction<CardToBoardBlindActionFormat>('CardToBoardBlindAction')

interface CardToHandBlindActionFormat {
	Player: string;
	BoardSlot: number;
}
export const CardToHandBlindAction = createAction<CardToHandBlindActionFormat>('CardToHandBlindAction')

interface RevealCardActionFormat {
	Name: string;
	Card: Card;
}
export const RevealCardAction = createAction<RevealCardActionFormat>('RevealCardAction')

interface PlayerReadyActionFormat {
	Name: string;
}
export const PlayerReadyAction = createAction<PlayerReadyActionFormat>('PlayerReadyAction')

export type CardActionTypes = CardToBoardAction | CardToHandAction | CardToBoardBlindAction | CardToHandBlindAction | RevealCardAction | PlayerReadyAction
*/
