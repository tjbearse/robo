import React from 'react'
import ReactDOM from 'react-dom'
import useImage from 'use-image';
import { Image } from 'react-konva';


export function UrlImage(props : {image:string})  {
	let [image] = useImage(props.image);
	props = Object.assign({}, props, { image });
	return <Image {...props} />
}
