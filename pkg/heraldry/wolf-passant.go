package heraldry

import svg "github.com/ajstarks/svgo"

func renderWolfPassantToSvg(canvas *svg.SVG, hexcode string, lineColor string) {
	pathData0 := "M97.7,98.2c6.8-2.9,13.3-4,18.1-9.1c1.2-1.2,3.4-1,4.7-0.2c1.3,0.9-0.1,2.4-0.6,3.3c-1.5,2.9-2.5,6.2-3.8,9.2    c-1.1,2.5-0.8,5.4-4.3,6.2c3.6,3.5-1.1,5-1.8,7.4c-1,3.1-1.4,6.4,0,9.1c3.1,6.1,3.4,12.6,4.1,19.1c0.5,5,0.8,10,1,15    c0.2,6.2,6,7.5,10.4,8.8c4.7,1.4,9.8,1.6,14.7,1.8c13.3,0.8,26.6-0.8,39.9-0.8c20.5-0.1,40.3,3.6,59.3,11.5    c8.1,3.3,13.4,0.9,18.5-6.2c3.5-4.9,0.2-15-5.5-17c-4.2-1.5-8.6-2.3-13.2-2.7c-6.4-0.5-12.6-2.3-18.3-5.1    c-6.8-3.4-14.3-5.2-19.8-11.2c-6.8-7.5-7-17.6-0.2-25.3c6.4-7.3,17-9.3,25.1-1.6c1.6,1.5,3.8,2.2,5,5.1c-4.4,0.5-9.6-1.2-11.8,4.4    c-0.9,2.3,0.6,3.8,1.6,6.1c0.9-2.4,1.4-4.5,3.2-5.8c1.4-1,2.1-3.6,4.3-2.7c1.6,0.7,0.4,2.8,0.7,4.2c0.4,2.1,1.8,3.7,3.7,4    c4.2,0.5,7.4,2.7,10.2,6c1-3,0.9-6.4,4.8-6.8c2.7-0.3,4.5,0.3,4.1,3.6c-0.5,3.6,1.4,5.2,4.7,5.7c6.5,1.1,10.9,5.4,15.3,9.8    c1.4,1.4,1.7,1.9,3.8,1.4c4.1-1.1,5.1,0.2,3.7,4.1c-0.6,1.7-0.3,2.7,1.3,3.3c0.7,0.3,1.5,0.6,2.3,0.9c6.4,2.3,6.7,3.2,2,8    c-1.2,1.2-1.7,2.1-1.2,4.2c1.3,5.7,0,11.1-3.6,15.9c-4.2,5.6-7.7,11.7-15.3,13.6c-6.3,1.6-11.9-0.1-17.6-2.8c1.9,2.4,3,4.8,3.3,8    c0.3,2.6,2.6,4.8,5.9,4.7c1,0,2.7-1,3.1,1c0.3,1.5-0.9,2.1-2.2,2.6c-2.4,0.8-2.5,1.6-1.4,4.3c1.8,4.5,2.1,9.6,2.7,14.4    c0.2,1.6,0.6,3,1.7,4.1c1.2,1.2,2.2,3.7,4.5,0.9c0.6-0.7,2.4-0.8,3,0.1c0.7,1-0.2,2.1-1.1,2.9c-5.3,5.1-5.3,5.1-2.4,10.4    c2.2,4.2,8,7,12.6,5.3c3.9-1.4,2.7,1.1,2.8,2.8c0,0.4-0.4,1.2-0.8,1.3c-6.4,2.1-4.5,7.2-4.2,11.6c0.1,1.7-0.8,2.9-1.5,4.2    c-5.7,10.7-11.2,21.5-17.8,31.6c-1.9,2.9-4.4,4-7.3,4.5c-3,0.5-5.8,1.6-8.2,3.3c-3.8,2.6-3.7,2.7-5-1.9c-0.2-0.6-0.6-1.2-1-2    c-1.8,1.1-3.2,2.8-5,4c-2.5,1.7-2.9,0.3-3.1-1.6c-0.7-4.9,3.6-12,8.2-13.8c2.1-0.8,4-2.4,5-4.3c5-10.4,8.2-21.9-3.7-31.9    c-6.2-5.2-12.8-10-19.1-15c-1.2-1-1.6-1.3-2.6,0.4c-4.1,7.3-3.1,13.5,3.5,18.2c3.1,2.2,0.4,4.3-0.7,4.4c-3.3,0.1-3.6,2.8-4.8,4.6    c-1.3,1.9-2.6,3.5-5,3.8c0.8,1.9,3.2,1.5,4.2,3.3c-2.2,1.4-4.4,2.9-7.2,2.3c-3.4-0.8-5.1,1.4-7.1,3.6c-7.5,8-16.2,14.5-26,19.3    c-2.3,1.1-4.8,1.7-6.5,3.9c-1.5,1.9-3.8,1.9-6.1,2c-4.3,0.1-9.1-1.4-12.4,3c-0.5,0.7-1.7,0.7-2.3,0c-1.6-1.9-2.9-3.1-5.8-2.4    c-1.5,0.3-1-1.9-0.6-3.2c1.7-6,5.2-9.4,11.7-10.2c8.7-1.1,13.4-8.2,19-13.9c0.7-0.7,1.2-1.8,2-2.6c3.7-4,4.1-8.4,1.9-13.2    c-3-6.7-4-13.9-5.1-21c-0.9-6,2.1-11.7,5.9-16.5c2.1-2.7,4.9-4.9,7.5-7.4c-1.9-1.1-4.1,0.4-5-1.1c-1.1-1.9,1.6-2.6,2.1-4.4    c-3.6-0.4-6.5,0.5-9.7,2.2c-12,6.7-24.6,12.1-38.5,13.7c-3,0.3-5.8,1.5-8.8,1.8c-3.7,0.4-4.3,3.1-4.4,5.9    c-0.2,5.1-1.6,9.5-4.7,13.6c-2.4,3.2-3.3,7.5-4.3,11.3c-2.7,10.2-7.8,19.4-11.3,29.1c-1,2.8-2.5,5.4-3.7,8.2c-1.3,3.2-4.1,4.5-7,5    c-3.7,0.6-7.1,1.7-10,4.3c-1.7,1.5-3.6,1.2-3.7-0.9c-0.3-4.6-2.5-1.5-4.1-1.4c-1.8,0.1-2.8,3.9-4.7,1.5c-1.5-1.8-1-4.4,0.1-6.7    c0.8-1.6,1.2-3.3,2-4.9c0.9-1.8,1.7-3.4,4.4-3.8c5.7-0.8,8.9-4.3,10.6-9.9c2.9-9.6,4.3-19.4,6-29.2c1.4-7.7,1.7-15.6,3.1-23.2    c0.4-2.1,0.2-3.3-2.8-3.6c-4.8-0.5-9.5-2.8-14.3,0c-4.8,2.8-9.4,6.3-14.8,7.6c-6.8,1.7-9.1,6.7-11,12.4c-0.6,1.9,0.3,3.2,1.7,4.3    c2.6,2.1,3.4,4.8,3.3,8.1c-0.1,4.2-1.1,8.8,3.1,12c0.6,0.5,0.5,1.8-0.3,2.2c-2.1,1.3-2.6,2.9-2.2,5.5c0.3,2.2-1.9,1.3-3.1,0.9    c-5.8-1.7-9.9-4.9-10.3-11.6c-0.2-2.5-1.1-5-2.6-7.1c-4.9-7-3.7-15.2-4.9-23c-1.3-8.3,1.9-13.9,9.4-17.9    c8.8-4.7,15.5-12.6,24.4-17.3c2-1,1.2-2.3,0.4-2.9c-3.7-3-2.5-8.9-6.9-11.4c-0.4-0.2-0.2-1.6-0.3-2.4c-0.2-2.4,0.1-4.7-1.3-7.2    c-1.7-3,0.9-5.7,3.7-8c-2.8-1.3-5.2-0.3-8.9-1.5c6.6-2,8.4-7.2,11.7-11.1c1.5-1.7,1.6-4,2.7-5.8c0.7-1.3-0.5-0.9-0.9-1.1    c-6.1-2.6-12.2-5-18.8-1.5c-2.3,1.2-4.5,2.8-6.2-0.9c-0.3-0.8-1.5-1.2-2.3-1.8c-1.6,0.7,0.3,4.2-2.9,3.5c-3-0.7-6.2-1-7.8-3.8    c-1.5-2.8,1.8-3.9,3-6.1c-3-0.9-5.7,1.1-8.5,0.7c-1.7-0.3-3.5,0-3.8-2.1c-0.2-1.8,1.3-2.5,2.8-3.1c4.9-2.2,10.2-2.9,15.6-3.6    c-0.4-2.4-1.8-3.1-4.1-3.1c-6.3-0.1-10.1-4.7-14-8.6c-1.8-1.8,0-4.1,3.5-4.9c5.7-1.3,11.4-2.8,17-4.1c7.2-1.7,13.6-4.6,19.7-8.4    c0.7-0.4,1.2-1,2.1-1.1c4.5-0.6,7.7-2.7,9.8-7c0.9-1.9,3.6-3.3,5.8-4.4c8-3.9,17-5.6,25-9.5c1.5-0.7,2.7-0.2,2.8,2.2    C102.5,91.6,100.1,94.4,97.7,98.2z"
	pathData1 := "M230,129.1c0.1-3.1,1.5,0,2.2-0.5c0.5,1.7,1.9,2.7,3.2,3.8c1.4-1.8-2.1-2.1-0.9-4.1c2.8,0.9,4.1,3.8,6.3,5.4    c2.3,1.7,3,4.7,5.6,6.2c1.4-1.7-2.2-2.1-0.9-4c2,2.5,4.6,4.4,6.1,7.3c0.1-1.9,0.5-3.3-1.6-5c-5-3.9-5.3-7.3-2-11.4    c1.2,0.6,1.3,1.6,1.2,2.8c-0.1,4.1,1.5,6.7,5.9,7.2c2.6,0.3,4.2,1.8,4.5,4.8c0.3,2.8-3.5,1.8-3.4,4.2c6-1.6,6.6-2.5,4.7-6.7    c2.8,2.4,6.1,4.4,8.1,7.2c2.6,3.6,4.6,5.5,8.2-0.6c-2,7.3-2,7.3,7.2,11.6c-4,5.1-4,5.1-10.6,0c-0.7,3,1.9,4.8,3.4,5.1    c2.9,0.7,3.1,2.5,3.6,4.7c1.6,6.2-2.2,13.1-8.2,15.4c-4.2,1.6-8.6,1.8-13,2c4.2,2.3,8.6,1.2,12.9,0.9c-7,5.4-14.7,3-22.1,1.9    c1.5,1.9,3.8,2.7,7.9,2.8c3,0,6.1,0,9.1,0.4c-4.4,3.4-9,3.5-14.3,1.6c-6-2.1-11.4-5.7-17.7-7.1c-5.9-1.4-11.8-2.4-17.9-2.1    c2.8,1.9,6.2,1.9,9.3,2.1c9.2,0.7,15.4,5.5,20,13.2c1,1.7-0.2,3.6,1.4,5c2,1.8,3.5,4.3,7.1,4c-4,2.3-6.9,1.7-8.7-1.9    c-1.6-3.3-4.7-5.4-6.5-8.5c-0.6-1-1.8-1.3-2.5-0.5c-1,1.1,0.3,1.9,1.1,2.1c2.5,0.9,3.4,3.4,4.6,5.1c2.1,3,4.1,5.3,7.5,7.1    c1.9,1.1,3,4.8,3.3,7.4c0.4,4,1.7,7.8,2,11.8c0.2,3.2,3.6,4.2,6,6.4c-7,1.7-11.3-2-15.5-6.1c-0.8-0.8-1.3-3.3-2.9-1.8    c-1.8,1.6,1.4,1.4,1.8,2.9c0.4,1.7,2.7,2.9,4.5,4.1c3.5,2.2,7.3,3.7,9.3,8.1c3.3,7.3,8.8,11.7,18,9.7c-5.6,7.6-11.8,3.5-17.9,2.6    c2.4,1.5,4.7,3.5,7.3,4.3c4.1,1.3,5,3.8,4.3,7.5c-0.3-0.2-0.5-0.4-0.7-0.5c-1.1-1.8-1.8-5.1-4.3-3.9c-1.7,0.8,1.7,2.3,2,4.2    c0.3,1.9,2.1,3.5,0.2,5.4c-0.4-0.8-0.6-1.8-1.2-2.5c-0.6-0.7-1.2-2.2-2.7-1c-1.9,1.5,0.7,1.5,1,2.1c1.3,2.5,1.6,4.4,0.1,7.1    c-3.9,6.8-7.5,13.9-11,21c-2.2,4.6-6.4,5.7-10.7,6.5c-2.3,0.4-4.7,0.1-7.4,0.1c1.1-4.3,3.8-6.5,7.6-7.7c1.1-0.3,2.7-1.2,2.1-2.4    c-0.9-1.9-1.6,0.5-2.4,0.6c-3.6,0.6-6.2,2.9-8.6,5.4c-1.6,1.8-3.3,1.7-5.5,1.2c1.2-2.7,1.9-5.1,5.7-6.1c3.9-1,6.7-4.7,8.3-8.7    c3.2-8.3,3-16.6-0.5-24.8c-0.6-1.4-2-2.4-3.4-4.1c2.4,0.2,3,2,4.3,2.9c0.7,0.5,1.3,2.1,2.3,1.1c1-0.9,0.1-2-0.7-2.9    c-2.6-2.8-5.9-4.6-9.2-6.1c-8-3.6-15-8.7-21-14.9c-4.4-4.5-6.9-9.7-7.8-16.2c-1.5-10.8-6.8-19.9-15.8-26.4    c-0.8-0.6-1.5-1.2-2.3-1.8c-0.7-0.6-1.1-1.7-3.3-1.4c4.1,7.2,8.5,13.8,13.8,19.9c-5.1-2.8-7.8-8.3-12.4-11.5    c-0.6-2.8-3.1-4.1-4.9-5.7c-1.6-1.4-3.7,0.5-5.7,0.6c-3.5,0.2-3.5,1.1-6.2-1.6c-1.1,1.3,0.4,3.6-1.4,4.3c-1.7,0.6-3.2-0.5-4.4-1.8    c-1.3-1.3-1.9-0.1-2.6,0.7c-3.2-1-6.4-1.5-9.3,1.4c3.6,0.6,7.5,0.5,11,2.1c-0.1-1-2.4-1.7-0.6-3.1c4.5,1.2,9.5-0.7,13.9,2    c1.8,1.1,4.1-0.2,6.1-0.4c1.6-0.2,3.7-1.3,4.2,1.6c-3.1-0.5-2.1,1.2-1.2,2.5c1,1.3,2.3,2.4,3.4,3.7c-2.4,1.1-3-1.1-4.3-1.9    c-1.8-1-3.8-4.4-5.8-0.1c-0.2,0.3-0.3,0.6-0.5,0.3c-3.2-4.2-7.5,2.1-10.8-1.1c0,0-0.2,0.2-0.6,0.4c4.3,4,9.6,5.3,15,6.1    c4.3,0.6,4.6,0.9,2.1,4.5c-1.9,2.8-4.4,4.7-8.4,3.8c2.4-2.8,5.4-4.6,8.8-6.4c-2-1.9-4.5-1.4-6.3-1.6c-5.6-0.8-10.7-3.1-15.9-4.8    c-0.5-0.2-1.1-0.5-1.6-0.4c-1.2,0.1-3-0.7-3.4,0.6c-0.5,1.8,1.5,1.9,3.1,2c-1,1.1-2.7,0.5-3.1,2c2.4,1.6,4.8,3.2,7.5,4.2    c-0.3,1.8-1.7,1.6-2,1.2c-3-2.7-7.3-0.8-10.5-3.1c-2-1.5-2.2,2.1-4.1,2.5c-3.4,0.6-6.6,2.2-10.6,0.3c4,5.1,8.5,2.8,11.5,0.9    c3-2,5-1.4,7.7-0.5c1,0.3,2.2,0.2,4.3,0.3c-4.9,3.3-8.6,6.9-15.3,2.9c1.3,2.5,3.3,1.6,3.4,3.1c0.2,0.4,0.1,0.5-0.4,0.6    c-3.8,0.1-7.3-0.2-10.3-3.5c-2.1-2.3-5.2-4.2-9.3-2.8c5.4,2.3,8.9,7.3,16,7.7c-4.9,2.2-4.9,2.2-17.6-3.1c0.5,1.8,3,4.2,0.3,4.8    c-2,0.4-4.3-1.7-5.7-4.2c2-0.4,3.4-1.4,5.2-1.3c1.6,0.1,3.2,0,5.8,0c-3.5-2.7-7.1-1.5-9.9-3.5c-2.9-2.1-1.8-6.2-5-7.9    c-0.3,3.3,0,6.4,2.2,9.1c0.3,0.4,1.4,0.8,0.5,1.6c-3.1,2.8-1,4.5,2.7,6.1c-5-0.2-6.5-2-5.4-5.1c0.2-0.5,0.8-0.9,0-1.3    c-4.2-2.4-5.7-6.5-6.9-10.9c-2.7,7.4,2.2,12.2,6,18.3c-4.1-2.3-6.1-5.8-9.6-7.7c1.8,2.8,3.5,5.5,5.3,8.4c-2.8-0.1-4.4-0.8-4.9-3.6    c-0.4-1.9-3.5,0.8-4.2-1.9c-1,2,0.1,3.5,1.1,4.3c1.9,1.6,5,2.4,1.4,5.5c-0.4,0.4,1.3,3.2,1.9,4.9c0.8,2.3,2.4,4.4,1.5,8.3    c-1.5-4.7-2.6-8.5-6.7-10.1c2.2,1.9,3.1,5.3,3.2,7.1c0.3,3.3,1.1,6.9-0.8,10.8c-4.1,8.6-6.2,18.2-9.7,27.1    c-2.9,7.4-6.5,14.6-9.9,21.9c-1.6,3.4-5.3,3.2-8.4,3.8c-2.2,0.4-4.5,0.1-6.9,0.1c0.7-4.3,3.6-6.3,7.1-7.5c1.4-0.5,3.1-1.3,2.5-2.4    c-1-2.1-1.7,0.5-2.8,0.7c-3.3,0.5-5.8,2.5-7.9,4.9c-1.6,1.8-3.3,2.4-5.7,1.5c0.8-3.2,1.6-5.4,6-5.8c5-0.4,9.3-6.6,11-12.1    c3.5-11.4,4.1-23.4,6.6-35c1.1-5.1,1.2-10.4,1.8-15.6c0.1-0.9-0.5-2.3,0.4-2.7c1.3-0.6,1.8,0.7,2.2,1.8c0.4,0.9,1.5,0.3,2.3,0.6    c2.3,0.9,4.5,2.1,6.6,3.5c-2.2-1.3-4.2-2.8-6.7-3.3c-0.6-3,1.5-2.3,2.7-1.5c3.6,2.4,5.6-1,8-2.2c0.8-0.4-1.5-2.3-1.6-4    c0-0.5-1-1.7-2.2-1.7c-0.2,0-0.4,0.3-0.7,0.4c2.7,2.6-3,3.5-1,6c-3.8-3.4-3.4-8.4-6.4-13.6c-0.5,4.5-0.4,7.9-0.4,11.4    c-1.1-4.9-4-9.2-4.6-15.3c-2.6,3.7,0.3,7.4-2.3,10.3c1.1-4.2-2-7.5-1.6-11.8c-2.9,3.2-0.1,6.6-0.9,9.9c-2.3-3.2-3-6.4-1.5-10    c2.3-5.4,0.9-12,5.9-16.6c1.3-1.2,3-2.3,4.2-4.3c-1.3,0-2.4,0-3.5,0c0.7-3.4,4.7-3.1,6.1-5.6c-1.7,1.3-3.5,2.6-5.5,3.6    c-0.9,0.5-1.3,0.9-0.5,1.8c-5.9,3.7-7.3,9.9-8.2,16c-0.3,1.9-0.7,3.4-1.8,4.9c-1.7,2.4-2.2,5.2-1.6,8.3c-2.8-1.3-3-2.7-2.1-5.3    c1.6-4.5,2.6-9.3,5.5-13.3c-6.6,3.1-5.7,10-8.1,15.3c-0.1-3.9,1.9-7.5,0.8-11.8c-3.2,4-1.8,9-3.7,13c-0.3,0.6,2,2.5,3.2,3.6    c1.1,1,2.4,1.9,4.4,3.4c-5.1,0.8-8.6,2.1-11.6,4.5c-1.6-1.4,0.8-2.5-0.2-4.5c-2.8,7.7-9.3,9.4-15.8,11.8c-2.4,0.9-5,3.2-5.9,6.7    c-0.7,2.8-1,6-3.4,8.3c-1.3,1.2,0.2,1.1,1.1,1.6c6.3,3.1,8.3,12.7,3.6,17.2c-1.4,1.4-2.3,1.6-2.7-0.6c-0.6-3.4-2.2-6.4-3.4-9.5    c-0.7-1.8-1.6-3.3-4.1-2.7c2.7,3.3,4.2,7.1,5.3,10.9c0.7,2.3,1.5,4.3,3.4,5.9c0.7,0.6,1.5,1.1,0.4,2.2c-1.1,1.2-1.6,0.6-2.6-0.1    c-1.7-1.2-2-3.1-2.3-4.7c-0.6-3.1-1.2-6.2-3.3-8.6c-3.9-4.5-2.2-10.3-3.6-15.3c-0.8-3.2-1-6.5-1-9.9c0-5.4,1.9-9.4,6.7-12.2    c7.1-4.1,13.4-9.5,19.8-14.6c2.4-1.9,5.2-3.3,8.2-5.2c-0.8,4,2,6.5,3.2,9.8c-2.7-9,3.3-15.3,6.9-22.3c-3.7,2.8-6.6,6-7.9,10.8    c-1.8-3.3-0.3-5.7,1.5-7.7c1.8-2,3.2-4.2,4.1-6.8c-3.8,2.7-6.4,6.3-8.6,10.2c-3.7-2.9-4.1-10.6-0.2-13.5c-3.3-1.2-4-0.3-4.3,5.4    c-2.8-1.9-3-3.7-1.7-7.5c1.4-4.1,4-7.8,4.5-12.2c0.2-0.8,1-1.2,1.1-2.5c-0.2,1.2-1.1,1.7-1,2.6c-3.8,2.1-3.6,7-7.4,9.8    c3-8.8,7.7-15.7,15.3-20.7c-3.3,0.1-5.3,2.4-7.4,4.5c-2.1,2.1-4.4,3.9-7.5,5.7c3.7-6.1,8.2-10.5,13.6-13.9    c0.7-0.4,2.2-0.6,0.9-1.9c-1-1.1-1.7-0.3-2.4,0.4c-1,1.1-1.9,2.2-3,3.3c-0.2-3.7,1.2-5.5,4.3-5.1c2.2,0.3,3.8,0.6,6.1-0.7    c1.9-1.1,4.9-0.3,7.3,0.9c-1.2,2.5-4.8,3.5-4.1,7.5c3.7-2.2,5.7-7.1,10.1-8.5c4.1-1.3,8.2-2.4,12.8-2.4c-0.7,2.9-4.1,3.5-5.2,6.1    c6.6-1.6,10.2-6.7,9.5-12.9c-0.1-1.2-0.4-2.5,0.4-3.7c3.6-5,1.6-9.6-0.7-14.3c-3.3,2.2,0.9,4.8-0.8,7c-2.2-1.1-1-3.5-1.8-5.5    c-3.2,2-6,5.3-11.3,3.1c3.2,3.5,7.7,4.1,8.2,9.2c-2.7-3.4-5.6-4-9.2-2.4c3.3,2.3,1.9,5.5,1.6,8.5c-0.1,0.6-0.2,1.1-0.9,1.2    c-5.8,0.3-11.7,1.5-16.9-2.4c-2.3-1.7,2.2-2.5,0-4.1c2.9,2.2,4.6,5.9,9,5c1.7-0.3,2.7-1,3.7-2.3c-1-0.5-2-1-3-1.4    c-1.2-0.5-1.3-3.7-3.1-2c-0.8,0.7,0.9,2.8,2.4,3.7c-7.1-2.3-10.9-10.3-19.2-11.3c1.9,3.1,5.6,3.8,7.3,6.4c2.5,3.8-2.8,5.6-2.4,9.3    c6.1-3.3,10.5,1.1,15.8,3.6c-2,1.3-3.2,0-4.3-0.7c-2.2-1.4-4.5-1.1-6.8-1.3c-1.6-0.2-0.5,4-3.1,1.6c-1.4-1.3-3.1-2.2-4.7-3.2    c-0.3-0.2-1,0.1-1.4,0.2c-0.4,2,2.5,1,2,3c-2.9,0.6-5.8-0.7-8.4-1.9c-5.4-2.5-10.4-1.7-15.5,0.6c-1.2,0.6-1.9,1.3-3.3-0.3    c-3.2-3.6-4-3.3-7,0.9c-0.6,0.8-0.9,0.2-1.2,0c-1.4-0.7-3.3-0.8-3.8-2.8c-0.2-0.9,0.4-1.2,1.2-1.7c4.3-2.8,9.1-1.8,13.4-1.2    c7.7,1,15.3,0.6,22.8-0.5c1.6-0.2,3.4-1.3,3.9-3.3c0.5-2.2-0.7-3.7-2.2-5c-3.1-2.9-7.2-2.4-10.9-3.1c-3.8-0.8-7.7-1.5-11.6-1    c-6.1,0.9-11.7-1.5-17.4-2.7c-1.6-0.3-2.9-2.2-4.3-3.4c1.4-0.9-0.4-4.3,2.9-4c-2,0-4,0-6.5,0c3.6-3,7.7-2.3,11.2-3.4    c8.2-2.6,16.9-4.1,24.3-8.9c7.8-5,16.4-6.5,25.4-6.4c3.3,0,5.6,0.3,3.4,4.3c-0.7,1.2-1.2,2.5-0.5,3.9c0.4-0.2,0.9-0.4,1-0.7    c4.6-8.9,13.3-11.4,21.8-14.6c3.3-1.2,7-2.5,9.4-5.6c1.1,1.3-0.6,1.7-0.6,2.4c-0.2,2.2-5.2,2.1-2.6,5.8c0.5,0.7-3,2.1-2.9,4.4    c0.1,2.3-2,2.8-3.2,4c0.9,0.7,2.6-0.5,2.9,0.9c0.4,1.3-1.1,1.6-2,2.2c-0.8,0.6-1.8,1-1.5,2.8c0.4,1.9-2.2,1.7-3.7,2.4    c0.9,0.5,1.6,1.1,2.4,1.2c3,0.4,1.8,1.9,0.7,3c-1.2,1.2-3.3,2.2,0.1,3.2c-5.6,2.2-10.9,4.4-16.5,6c5.7,0.9,11,0,15.8-3.4    c1.8-1.2,3.3-2.1,4.3,1.4c1.5,5.1,1.6,10.3,2.2,15.4c0.6,5.4,1.4,10.8,1.1,16.3c-0.2,3.4,2.4,5.5,5.1,7c3.8,2.2,8,3,12.4,3.5    c10.9,1.3,21.8,0.5,32.8,0.7c7.9,0.2,15.7-1.4,23.6-0.9c7.6,0.5,15,1.9,22.5,2.9c12.1,1.7,23.1,7.2,34.9,9.6    c5.2,1.1,12.9-1.4,15.7-5.7c3.6-5.4,4.8-10.9,1.3-17.2c-2.5-4.4-6.3-6-10.5-7.3c-6.9-2.2-14.4-1.5-21.2-4.3    c-8.9-3.7-18.7-6.1-25.4-13.4c-8.3-9-3.3-21.5,6.3-26.3c0.3-0.1,0.8,0.1,1.2,0.2c-8.1,7.1-9.6,8.5-4.7,16.5    c-1.8-9.7,3.1-14.4,10.2-17.5c-2.6,3.2-6.6,5.1-7.5,9.9c-0.9,4.5-0.4,8.4,2.4,11.9c-0.6-2.9-0.3-5.7-0.3-8.6    c0-4.9,3.3-9.4,7.9-10.9c0.6-0.2,1.1-0.9,1.7,0c0.8,1.2-0.4,1.3-0.9,1.8c-7.3,5.3-8.6,9-6.1,17c0.4,1.2,1.3,2,3,2.7    c-1.1-4-2.7-7.9-1.9-11.4c0.9-3.6,3.4-7.1,7.8-8c-3.3,2.7-6.9,5.3-6.3,10c0.7,5.5,5.3,11.3,10.1,13.3c-0.5-3-4.6-3.9-3.9-8.4    c1.8,2.1,3.1,3.7,4.4,5.4c1.3,1.6,2.2,3.6,4.3,4.8c-0.5-4-3.8-5.9-5.8-8.7c-1.9-2.7-2.2-4.8,0.3-7.2c1.2,2.6,0.9,5.9,4.1,7.2    c0.4,2,2.5,2.9,3.2,4C232.5,131.9,230.8,130.9,230,129.1z"
	pathData2 := "M216.6,265.2c-4.1,0.7-9.9-3-11.1-7.6c-1.5-6.1-6.3-8.9-10.5-12.4c-0.3-0.2-0.7-0.2-2.3-0.6    c3.4,4.5,7.9,6.6,10.4,10.9c-0.7-0.1-0.9,0-1.1-0.1c-1.5-0.7-2.3-4-4.4-1.9c-0.8,0.9,1.7,2.8,2.9,3.7c3.6,2.7,5.3,6.9,8.5,9.9    c1.6,1.5,0.8,4.2-1.8,5.5c-1.4-9-14.9-23.7-23.1-25c7.7,8.2,19.2,13.2,21.4,26.2c-3-0.7-3.7-2.7-4.4-4.6c-1.4-4-3.4-7.7-6.4-10.8    c-0.5-0.5-1-1.2-1.7-0.6c-0.3,0.3-0.5,1.3-0.3,1.5c5,4.2,5.6,11.1,9.6,15.9c1,1.2,2.4,2.2,3.6,3.3c-4,2.9-7.4-0.1-8.1-5.1    c-0.6-3.5-1.5-6.8-4.7-9c0.9,4.4,4.1,8.3,2.7,13.4c-0.8-1.8-2.7-2.7-2.3-6c-1.4,3.3-0.9,5.4,0.4,6.7c2.1,2.1,0.9,3.6-0.4,4.8    c-7.1,6.9-14.8,13.1-23.4,18.1c-0.5,0.3-1.6,2.1-2.7,0.2c-1.1-2-1.1-0.8-1.8,0.5c-3.6,7.4-13,8.5-18.4,2.2    c4.6-3,11.4-2.1,14.3-7.9c-4.2,2-8.3,4.3-12.9,5.1c-2,0.3-2.7,1.9-3.9,2.9c-0.9,0.8-1.3,1.5-2.5,0.3c-1.4-1.4,0-1.7,0.5-2.5    c1-1.6,2.9-2,4.3-2.1c10.6-1.5,16.6-9.1,23-16.3c0.4-0.4,0.6-1,1-1.3c6.6-4.7,5.6-10.8,2.9-17c-2.6-6.1-3.1-12.6-4.7-19    c-0.3-1.1-0.1-2.3,0-3.4c0.6-5.5,2-6.1,6.7-3.3c6.3,3.7,13,6.9,18.4,12c1.1,1,2.2,0.8,3.5,0.4c-6.4-7.8-16.1-10.8-23.4-17.2    c4.9,0.3,8.3,3.7,12.3,5.9c1.5,0.9,2.9,2.1,4.2,3.3c0.8,0.7,1.4,0.8,2.1,0c1-1.1-0.1-1.3-0.6-1.8c-4.4-4-9.5-6.9-14.8-9.5    c-0.8-0.4-3.4-0.6-2.1-2c1.2-1.3,2.4-3.7,4.7-3.3c1.5,0.3,3.1,1.1,4.4,2.1c6.1,4.6,12.8,8.6,17.7,14.7c0.8,1,1.5,1.7,2.8,1.2    c-2.5-4.6-4.5-6.3-10.6-11.5c-3.9-3.4-9-5.1-12.8-9.5c8.3,2,13,8.6,19.7,12.3c-3.4-6.4-10.4-9.2-15.7-14.7    c4.4,0.7,6.8,3.3,9.3,5.7c0.6,0.6,1.1,2.4,2.4,1.2c1.3-1.3-0.6-1.8-1.2-2.4c-1.1-1.1-2.4-1.9-3.4-3.1c3.6-0.1,9.4,4.9,9.9,7.7    c-1.9,0.9-2.5-3.5-4.1-1.4c-0.9,1.3,1.8,2.5,2.9,3.8c2.1,2.4,5,4.2,6.8,7.4c-2.7,0.2-2.9-3.1-5-3.1c-0.5,0,0-2.4-1.1-0.4    c-0.3,0.5-1.3,0.9,0.3,1.6c2.5,1.2,3.6,4.2,5.6,6c1.4,1.3,0.6,2.4,0,3.2c-1,1.2-1.3-0.5-2-0.8c-7.4-3.4-12.1-10.4-19.2-14    c-2.8-1.5-5.4-5.5-10-2.4c6.8,1.2,11.2,6.1,16.2,9.7c4.7,3.3,9.6,6.7,13.3,11.2c0.8,1,3.3,2.6-0.3,3.4    C209.8,257.9,211,263.3,216.6,265.2z"
	pathData3 := "M96.1,93.9c-0.3,1.7-2,1.1-3,1.9c0.7,0.2,1.3,0.4,2,0.6c-0.9,1.4-2.6-0.1-3.5,1c0.7,0.9,1.9,0.1,2.8,1.3    c-5.9-0.5-11.5,0.7-17.4,2.2c5.1,1.1,9.9,0.5,15.4-0.9c-3.3,3.8-6.8,3.4-10.2,2.6c-2.8-0.7-5.6-0.6-8.3-0.2    c-1,0.1-2.2,1.3-2.7-0.1c-0.3-0.7,0.5-2,1.7-2.8c4.2-2.9,8.5-5.4,13.2-7c4.1-1.4,8.2-2.5,12.4-3.8c-0.5,2.9-1.7,4.8-5,4.2    c-1.4-0.3-2,1.2-2.9,2.1c1.5,1.1,1.8-0.8,2.7-1.1C94.2,93.5,95.2,93.6,96.1,93.9z"
	pathData4 := "M48.2,142.7c-1.3,0-2.4-0.3-3.3,0.1c-1.7,0.7-3,0-3.1-1.3c-0.1-1.1,1.3-2.2,2.9-2.2c5.7,0,11.3-0.4,16.6,2.4    c0.9,0.5,2.4,0.1,2,1.8c-0.5,1.6-1.9,1-2.8,0.7C56.4,143.2,52.2,142.5,48.2,142.7z"
	pathData5 := "M19.8,143.3c3.8-3.7,8.8-3.4,13.3-4.8c1.2-0.4,3.5-0.6,3.6,1c0.1,1.4-1.5,3.1-3.3,3.1    C28.9,142.8,24.3,142.5,19.8,143.3z"
	pathData6 := "M63,140.1c-5.5-3.5-11.6-2.2-17.4-2.5c-0.9,0-2.5,0.7-2.4-1.3c0-1.9,1.6-1.1,2.5-1.2c5.1-0.4,10,1.9,15.1,1.7    C62.5,136.8,62.4,138.7,63,140.1z"
	pathData7 := "M36.7,135.4c1.4,0.3,3.3-1.5,3.8,1c0.1,0.6-0.1,1.9-0.3,1.9C38.2,138.4,37.8,136.5,36.7,135.4z"
	pathData8 := "M144.3,311.8c-0.6-0.6-1.3-1-1.3-1.4c-0.1-1.8,1.6-2.5,2.7-3.4c0.6-0.5,1.4,0.6,1.3,1.1    C146.7,309.7,144.1,309.6,144.3,311.8z"
	pathData9 := "M51.7,275.6c-0.6,0.6-1,1.3-1.4,1.3c-1.8,0.1-2.5-1.6-3.4-2.7c-0.5-0.6,0.6-1.4,1.1-1.3    C49.5,273.2,49.4,275.8,51.7,275.6z"
	pathData10 := "M165.8,214.4c2.2-3.4,4.3-0.3,6.3,0C170,214.4,168,214.4,165.8,214.4z"
	pathData11 := "M48.8,281.8c-1-0.9-3.4-0.4-3-2.6c0.1-0.4,1-1.2,1.1-1.1C48.1,278.9,48.6,280.2,48.8,281.8z"
	pathData12 := "M138.2,309c0.5-1.3,0.4-3.2,2.3-3.2c0.4,0,1.2,0.8,1.1,1C141.3,308.6,139.5,308.5,138.2,309z"
	pathData13 := "M37.8,140.8c1.9,1,1.9,2.2,2.5,3.2c-1.2,0.7-2.4,0.3-3.6,0.3C36.6,142.9,38.7,142.7,37.8,140.8z"
	pathData14 := "M231.6,305.2c-1.1,1.1-2,2.1-3.3,3.5C228.1,305.7,228.8,304.6,231.6,305.2z"
	pathData15 := "M66.8,306c-1.1,1.1-2,2.1-3.3,3.5C63.3,306.6,64,305.5,66.8,306z"
	pathData16 := "M75.7,308.4c-1.1,1.1-2.3,1.8-3.7,2c-0.1-0.4-0.3-0.9-0.1-1.2C72.8,307.2,74.5,309,75.7,308.4z"
	pathData17 := "M240.5,307.8c-1.2,0.8-2.2,1.9-3.7,1.7c-0.1,0-0.2-0.9-0.1-1.2C237.7,306.6,239.2,307.8,240.5,307.8z"
	pathData18 := "M46.2,144.6c1.9,0,3.8,0,5.7,0c0,0.1,0,0.3,0,0.4c-1.9,0-3.9,0-5.8,0C46.1,144.8,46.2,144.7,46.2,144.6z"
	pathData19 := "M162.3,210.3c1.8,0.4,3.3,0.8,5.7,1.4C165.4,211.9,163.7,212.8,162.3,210.3z"
	pathData20 := "M98.8,108.2c2.1,1.6,3.6-1.5,5.3,0.6c-1.7,2.3-5,1.4-7.6,3.8c2.4-0.2,4-0.3,5.2-0.3c-0.4,5.1-4.2,4.3-7.5,4.3    c0.3,2.4,2.4,1.3,3.8,2c-1.9,1.2-4.4-0.6-5.4,1.6c-1.3,2.7-3.4,0.9-5.1,1.3c3.5-2.5,3.4-6.1,3-9.6c-0.4-3.6,1.3-5.6,4.5-6.3    c4.1-0.9,6.2-5.4,10.7-5.6c0.9,0,1.9-0.9,2.5,0.2c0.4,0.7-0.2,1.5-0.6,2.3C105.8,106,102.8,107.5,98.8,108.2z"
	pathData21 := "M71.3,117.5c1.9,2.9,2.6,5.6-1.3,7.5c-1.2,0.6-1.4,1.3-0.7,2.5c0.3,0.5,1.4,1.2,0.1,2.1    c-1.3,0.9-2.8,0.7-4,0.3c-1.1-0.4-1.1-1.2,0.2-1.8c0.3-0.2-0.9-0.9-1.5-1.4c-1.7-1.1-4.1-0.8-5.7-3c5.2,0.8,0.3-4.8,4.4-4.9    c-2.9-0.8-5.5,0.2-7.7-1.9c1.8-1.3,3.5-2.5,5.8-4.2c-0.6,1.7-1,2.6-1.4,3.9c7.9,0.3,15.4-1.5,23.4-2.7    C79.9,118.9,75.2,116.8,71.3,117.5z"
	pathData22 := "M187.9,202.1c-2.3-2.7-4.5,0.1-6.7,0.1c-0.8,0-2.1,0.9-2.7,0.4c-4.6-3.5-10.2-1.7-15.2-3.3    c-0.8,0.2-1.2-0.1-1-1c0.4-1-0.8-2.7,0.9-3.2c1.4-0.4,2.1,0.6,2.7,2c0.6,1.5,2.3,1.7,3.8,2.3c-0.1-2.5-0.2-4.6-0.3-7    c3,1.3,3.6,5.4,7.4,5.6c-0.3-1.5-2.9-2-1.5-3.3c1.5-1.6,2.1,0.3,2.9,1.3c0.7,0.8,1.8,1.4,1.9-0.4c0.1-1.5,0.9-1.5,1.9-1.4    c1.5,0.1,7.1,6,7,7.4c0,0.1-0.7,0.2-1.1,0.3C187.8,201.9,187.9,202.1,187.9,202.1z"
	pathData23 := "M162.2,198.4c0.3,0.3,0.7,0.7,1,1c0.2,1.8,2.3,1.4,2.9,2.6c0.3,0.6,1.7,0.8,0.5,1.9c-0.8,0.7-1.3,0.6-2.1-0.1    c-3-2.4-6.2-2.5-10-2.1c-2.8,0.3-3.3-2.3-3-4.8c2-0.3,1.2,2.2,2.7,2.1c0-2,3.7-4.2,5.2-3C160.4,196.8,160.9,198,162.2,198.4z"
	pathData24 := "M43.8,242.3c-1.8,0.9-3.3,2.1-4.6,3.6c-0.5,0.5-1,1.1-1.7,0.5c-0.3-0.3-0.5-1.3-0.2-1.5    c2.6-2.2-0.1-6.5,3.3-8.4c0.1-0.1,0-0.5,0.1-0.8c2.8,1.5,5.2-0.3,7.7-1c0.5-0.1,0.5,0.8-0.4,1.1c-1.7,0.5-2.7,2.3-4.6,2.5    c-1.5,0.1-1.6,1.6-1.7,2.4C41.3,242.2,43.2,241,43.8,242.3z"
	pathData25 := "M254.6,229c-4.7-3-7.6-7.3-12.8-8.1c-0.9-0.1-0.9-0.6-0.1-1.3c0.2-0.2,1,0.1,0.6-0.7c-0.3-0.6-0.7-1.1-1-1.7    c0.1,0,0.3-0.1,0.4-0.1c0.9,0.2,1.8-0.4,2.8,0.9c1.8,2.1,4,4.7,6.2,5.5C253.5,224.6,254,226.1,254.6,229z"
	pathData26 := "M259.1,275.6c0.8,1.5,0.1,2.7-0.4,4c-0.3,0.7-0.3,1.8,1.1,1.6c0.8-0.1,1.7-1.7,2.1,0.2    c0.2,0.8-0.8,1.3-1.6,1.6c-3.1,1.4-2.8,4.7-4.1,7.2c-0.4,0.7,0.4,2.1-1.3,2.1c-0.8,0-1.5-0.2-1.8-0.9c-0.6-1.2,0.8-1.6,1.3-2.2    c2.6-3,1.6-7,3.1-10.3C258.2,277.8,258.4,276.7,259.1,275.6z"
	pathData27 := "M99,247c1.3-0.7,3.4,0.3,3.7-1.4c0.3-1.3-1.4-2.3-2.5-2.9c-1.3-0.7-1.9-1.5-1-2.5c1-1.1,1.7,0.4,2.2,1    c1.9,2.1,5.6,3.8,1.9,7.4c-1.2,1.2,2.7,3.4-0.5,3.9c-2.7,0.4-0.5-3.2-2.4-3.9C99.8,248.3,99.5,247.5,99,247z"
	pathData28 := "M245.2,209.3c5,2.3,5.4,7.7,8.4,11.4c-5-2.3-7.4-6.3-8.3-11.5L245.2,209.3z"
	pathData29 := "M82.6,164.1c-2.1,4.2-5.8,6.6-10.6,8.8C78.3,164.5,79,163.9,82.6,164.1z"
	pathData30 := "M235.9,208.1c2.3-0.1,8.4,5.1,8.3,7.1c0,0.3-0.3,0.9-0.5,1c-1.2,0.5-2-0.7-3.1-0.9c0.9-4-3.9-3.7-4.7-6.4    C235.8,208.7,235.9,208.6,235.9,208.1z"
	pathData31 := "M245.3,209.2c-4.1,0.2-6.8-2.5-10.7-5.2c3.6-0.2,4.7,2.4,7,2.6c0.2-2.5-3.2-2.7-3-5l-0.1,0.1    c4.1,0.9,4.4,5.2,6.7,7.6C245.2,209.3,245.3,209.2,245.3,209.2z"
	pathData32 := "M267.9,150.5c-0.1,2.2-1.2,4.2-0.9,6.5c0.1,0.7-0.4,1.7-1.5,1.3c-0.8-0.3-1.1-0.9-0.7-2    c1-2.7,2.8-5.6-0.2-8.3c-0.1-0.1,0.3-1,0.6-1.4c0.8-1,1.1,0.1,1.6,0.4C268,147.9,268.1,149.1,267.9,150.5z"
	pathData33 := "M89.5,191.8c-0.7-0.1-1.5,0-2-0.3c-0.5-0.5,0.3-0.9,0.6-1.2c3.1-2.2,6.1-4.3,9.6-6.6c-0.6,5.5-6,5.2-8.3,8.2    C89.4,191.9,89.5,191.8,89.5,191.8z"
	pathData34 := "M254.3,232.6c-4-3.6-9.8-5-12-10.5C246.3,225.5,251.1,228.1,254.3,232.6z"
	pathData35 := "M256.1,246.4c-3-3.8-7.4-4.6-8.9-8.6C250.6,237.6,255.2,241.5,256.1,246.4z"
	pathData36 := "M106.7,173.8c-2.2-0.2-3.9,1.6-6.8,2.1c3.8-4.9,9-4.7,13.7-7.6c-1.9,3.6-5.9,2.8-7,5.7L106.7,173.8z"
	pathData37 := "M106.4,153.9c-0.5,5.4-4.4,8.7-10.2,8.9C100.3,160.6,103.5,158,106.4,153.9z"
	pathData38 := "M220.9,234.3c1.7,3.2,6.1,4.6,6.6,9.5C223.1,241.4,222.1,237.8,220.9,234.3z"
	pathData39 := "M213.2,223.6c3,3.3,2.7,6.9,4.4,10.6C214.1,231.9,212.9,228.7,213.2,223.6z"
	pathData40 := "M74.1,186.9c-1.1,2.8-4,4.2-5,7c-0.2,0.6-1.3,0.5-1.9,0.2c-0.9-0.6-0.2-1.1,0.2-1.7    C68.9,189.6,71.1,187.8,74.1,186.9z"
	pathData41 := "M262.7,253.6c-4.1-0.8-6.5-3.7-8.7-6.9c3.1,2.1,7.1,2.9,8.6,7L262.7,253.6z"
	pathData42 := "M106.6,173.9c3.7-0.8,6.8-3.5,10.9-3.4c-4.1,1.9-8.3,3.4-11.7,6.6c-0.2-0.3-0.4-0.6-0.6-0.9    c0.5-0.8,1.1-1.6,1.6-2.5C106.7,173.8,106.6,173.9,106.6,173.9z"
	pathData43 := "M276.9,176.8c-1.9,4-5.4,4.7-9.5,4.3C270.5,179.6,273.7,178.2,276.9,176.8z"
	pathData44 := "M78.2,171.6c-2,2.9-5.2,4.2-8.2,6.9C71,173.3,75.1,173.1,78.2,171.6z"
	pathData45 := "M75.2,181.1c-1.2,3.8-3.4,5.4-6.6,6C70.5,185.2,71.3,182.2,75.2,181.1z"
	pathData46 := "M233.8,189.3c-4.8-0.5-8.5-2.3-12.5-2.2C225.3,187.2,229.4,185.3,233.8,189.3z"
	pathData47 := "M253.5,265.6c-3.5-1.6-2.2-5.1-3.6-8.1C253.7,259.6,252.6,262.9,253.5,265.6z"
	pathData48 := "M208.4,216.5c1.4,2.7,3.2,5.3,2.7,9.9C209.3,222.5,207.7,219.9,208.4,216.5z"
	pathData49 := "M95.3,235c4.2-0.9,5.9,3.4,9.2,4.4c-3.5-0.6-6.4-2.3-9-4.5C95.4,234.9,95.3,235,95.3,235z"
	pathData50 := "M254.2,266.1c1-2.9,1.2-5.5-0.1-8.4C257.9,260.5,257.9,263.1,254.2,266.1z"
	pathData51 := "M268,177.8c4-0.9,5.6-4.1,9.5-4.8C274.8,175.8,273.1,178.7,268,177.8z"
	pathData52 := "M270.5,157.6c-0.4,1.2-1.1,2-1,3.2c0.1,1.3-0.7,2-2,2.1c-0.4,0-1-0.1-1.1-0.3c-0.6-1.2,0.5-1.7,1.2-2.3    c0.9-0.8,1.3-1.7,1.1-2.9c-0.1-0.8,0.4-1.4,1.2-1.1C270.3,156.4,270.4,157.2,270.5,157.6z"
	pathData53 := "M273.9,167.3c-0.7,1.4-2.7,1.7-3.3,3.4c-0.2,0.6-1.3,1.2-1.6,0.4c-1.3-2.6,1.6-3.7,2.6-5.1    c1.4-2,0,1.4,1.2,1.2c0.4-0.1,0.8,0.2,1.2,0.2L273.9,167.3z"
	pathData54 := "M93.1,142.2c-3.4,2.9-4.9,0.2-7.4-2.2C88.6,140.9,90.6,141.5,93.1,142.2z"
	pathData55 := "M261.5,277.9c-0.3-2.5-0.6-5-0.9-7.5C262.1,272.7,263.6,275.1,261.5,277.9z"
	pathData56 := "M50.3,235.3c1.1-3.1,2.8-4.5,4.5-6.3C55.2,232.2,53.7,233.8,50.3,235.3z"
	pathData57 := "M249,262.5c-1.4-0.4-1.9-2-2-2.8c-0.2-1.4-3.2-2.1-1.2-3.5c0.8-0.5,1.8,1.9,2.6,3    C249,260.1,250,261.2,249,262.5z"
	pathData58 := "M258.7,257.7c0.8,2.3,2.1,4.7,0.4,7.1c-1.2-1,0.1-2.5-0.8-3.5C257.4,260,257.6,258.8,258.7,257.7z"
	pathData59 := "M97.9,167.1c1.9-1.1,3.7-2.2,6.4-3.8c-1.2,3.9-4,3.5-5.7,4.9C98.3,167.8,98.1,167.5,97.9,167.1z"
	pathData60 := "M230,129.1c2.2,2.1,4.4,4.2,6.6,6.3c-2.9-1.4-6.1-2.6-6.5-6.4C230.1,129,230,129.1,230,129.1z"
	pathData61 := "M76.3,204.1c-1.4-2.9,0.4-4.3,2.9-5.9C79.7,201.1,75.9,201.6,76.3,204.1z"
	pathData62 := "M75.4,175.9c-0.4,4.3-2.5,5-5.6,5.8C71.8,179.6,73.3,178.1,75.4,175.9z"
	pathData63 := "M237.2,193.1c-2.9-0.5-5.6-1.7-8.5-1.7C231.8,190.5,234.7,190.8,237.2,193.1z"
	pathData64 := "M81.7,136.2c3-0.6,4.8-0.2,6.9,1.8C85.7,138.4,84,137.2,81.7,136.2z"
	pathData65 := "M206.2,178.9c-1.7,0-3.2,0-5.5,0C203,177.1,204.5,176.4,206.2,178.9z"
	pathData66 := "M114.2,174c1.8-0.7,3.3-2.3,5.5-1.5C118.3,174.6,116,173.5,114.2,174z"
	pathData67 := "M100.1,171.4c-0.5-0.1-1.2,0-1.1-0.8c0.1-1.1,0.8-1.7,1.9-1.7c0.7-0.1,1.5,0,1.3,1    C102,170.8,101.2,171.5,100.1,171.4z"
	pathData68 := "M195.7,175.3c1.6-0.3,3-1.9,5-1.1C199.4,176,197.6,175.8,195.7,175.3z"
	pathData69 := "M64.5,173.3c-0.7-2.8,2.2-2.9,3.6-4.2c-1.3,1.4-2.5,2.7-3.8,4.1C64.3,173.2,64.5,173.3,64.5,173.3z"
	pathData70 := "M213.6,179.2c-1.6,1.7-3.4,1.6-6.1,1.2C210,179.8,211.5,178.5,213.6,179.2z"
	pathData71 := "M271.4,173.2c1.3-0.7,2-2.2,3.7-2.9C274.8,172.9,272.9,172.8,271.4,173.2z"
	pathData72 := "M275.9,169.7c-0.8-1.9,0.3-2.8,0.9-3.9c0.2,0.2,0.6,0.3,0.7,0.5C278,167.7,277.1,168.7,275.9,169.7z"
	pathData73 := "M262.5,253.7c1.7-0.8,2.9,0.1,4.6,1.1c-2.2,0.6-3.3-0.5-4.5-1.2C262.7,253.6,262.5,253.7,262.5,253.7z"
	pathData74 := "M274.1,167.4c-0.1-1.1-0.2-2.1-0.3-3.5c2.8,1.3,1.6,2.3,0.2,3.4C273.9,167.3,274.1,167.4,274.1,167.4z"
	pathData75 := "M143.6,201c0,0.6-0.4,0.9-1,0.8c-0.5-0.1-1.3-0.2-1.4-0.5c-0.3-0.9,0.4-1.1,1.2-1.1    C143,200.2,143.5,200.2,143.6,201z"
	pathData76 := "M86,132.7c1,0,2.1,0,3.8,0C88,133.7,86.8,133.3,86,132.7z"
	pathData77 := "M102.3,255.8c-0.2,1.1-0.9,0.9-1.5,0.9c-0.3,0-0.9,0.1-0.6-0.6c0.3-0.6,0.7-1.2,1.5-1    C101.9,255.2,102.1,255.6,102.3,255.8z"
	pathData78 := "M262,147.4c-0.2,0.2,0.3,1.2-0.6,0.8c-0.4-0.2-0.8-0.7-1-1.2c-0.2-0.6,0.2-1,0.9-0.9    C262,146.2,262,146.8,262,147.4z"
	pathData79 := "M257.8,139.3c-0.2,0.8-0.7,1-1.2,0.7c-0.4-0.2-0.7-0.7-1.1-1.1c0.4-0.1,0.9-0.4,1.3-0.4    C257.5,138.4,257.7,138.8,257.8,139.3z"
	pathData80 := "M238.6,201.6c-0.9,0.2-1.9,0.7-2.6-0.3c-0.1-0.2,0.1-1,0.3-1c1.2-0.4,1.6,0.7,2.2,1.4    C238.5,201.7,238.6,201.6,238.6,201.6z"
	pathData81 := "M181.9,247.6c-2.4-1.1-4.8-2.1-7.1-3.4c-0.5-0.3-1.7-0.7-1.1-1.8c0.6-1,1.3-0.8,2.4-0.4    C178.8,243.1,180.2,245.6,181.9,247.6z"
	pathData82 := "M179.9,254.6c1.3,2,2.4,3.6,3.3,5.3c0.3,0.6,0.6,1.7-0.7,2c-0.8,0.2-1.7,0.1-1.7-1.1    C180.9,258.7,178.2,257.4,179.9,254.6z"
	pathData83 := "M183.6,253c-3.2,1.2-3.7-2.2-5.8-2.4c-0.1,0-0.2-1-0.1-1c0.5-0.2,1.2-0.6,1.5-0.3    C180.6,250.3,181.9,251.5,183.6,253z"
	pathData84 := "M177.9,266.2c1.6,0.6,3.1,1.3,2.9,3.3c0,0.3-0.6,0.7-0.8,0.7C178.5,269.5,178.6,267.6,177.9,266.2z"
	pathData85 := "M82.1,97.5c1.2-0.6,2.2-1.7,4.1-1.1C84.7,97.2,83.8,98.5,82.1,97.5z"
	pathData86 := "M98.8,115c-0.3,0.9-1.1,0.7-1.7,0.6c-0.7-0.1-1-0.7-0.7-1.1c0.2-0.3,0.9-0.4,1.4-0.5    C98.6,113.9,98.8,114.3,98.8,115z"
	pathData87 := "M93.8,110.7c0.8-1.2,0.3-2.5,2.5-2.8C95.6,109,96,110.3,93.8,110.7z"
	pathData88 := "M69,118.7c-2.7-0.2-3.1,1.7-3.4,3.6c-1.4-0.8-1.8-1.8-1.2-2.9C65.5,117.5,67.3,118.2,69,118.7z"
	canvas.Translate(30, 30)
	canvas.Scale(0.8)
	canvas.Path(pathData0, "fill:"+lineColor)
	canvas.Path(pathData1, "fill:"+hexcode)
	canvas.Path(pathData2, "fill:"+hexcode)
	canvas.Path(pathData3, "fill:"+hexcode)
	canvas.Path(pathData4, "fill:"+hexcode)
	canvas.Path(pathData5, "fill:"+hexcode)
	canvas.Path(pathData6, "fill:"+hexcode)
	canvas.Path(pathData7, "fill:"+hexcode)
	canvas.Path(pathData8, "fill:"+hexcode)
	canvas.Path(pathData9, "fill:"+hexcode)
	canvas.Path(pathData10, "fill:"+hexcode)
	canvas.Path(pathData11, "fill:"+hexcode)
	canvas.Path(pathData12, "fill:"+hexcode)
	canvas.Path(pathData13, "fill:"+hexcode)
	canvas.Path(pathData14, "fill:"+hexcode)
	canvas.Path(pathData15, "fill:"+hexcode)
	canvas.Path(pathData16, "fill:"+hexcode)
	canvas.Path(pathData17, "fill:"+hexcode)
	canvas.Path(pathData18, "fill:"+hexcode)
	canvas.Path(pathData19, "fill:"+hexcode)
	canvas.Path(pathData20, "fill:"+lineColor)
	canvas.Path(pathData21, "fill:"+lineColor)
	canvas.Path(pathData22, "fill:"+lineColor)
	canvas.Path(pathData23, "fill:"+lineColor)
	canvas.Path(pathData24, "fill:"+lineColor)
	canvas.Path(pathData25, "fill:"+lineColor)
	canvas.Path(pathData26, "fill:"+lineColor)
	canvas.Path(pathData27, "fill:"+lineColor)
	canvas.Path(pathData28, "fill:"+lineColor)
	canvas.Path(pathData29, "fill:"+lineColor)
	canvas.Path(pathData30, "fill:"+lineColor)
	canvas.Path(pathData31, "fill:"+lineColor)
	canvas.Path(pathData32, "fill:"+lineColor)
	canvas.Path(pathData33, "fill:"+lineColor)
	canvas.Path(pathData34, "fill:"+lineColor)
	canvas.Path(pathData35, "fill:"+lineColor)
	canvas.Path(pathData36, "fill:"+lineColor)
	canvas.Path(pathData37, "fill:"+lineColor)
	canvas.Path(pathData38, "fill:"+lineColor)
	canvas.Path(pathData39, "fill:"+lineColor)
	canvas.Path(pathData40, "fill:"+lineColor)
	canvas.Path(pathData41, "fill:"+lineColor)
	canvas.Path(pathData42, "fill:"+lineColor)
	canvas.Path(pathData43, "fill:"+lineColor)
	canvas.Path(pathData44, "fill:"+lineColor)
	canvas.Path(pathData45, "fill:"+lineColor)
	canvas.Path(pathData46, "fill:"+lineColor)
	canvas.Path(pathData47, "fill:"+lineColor)
	canvas.Path(pathData48, "fill:"+lineColor)
	canvas.Path(pathData49, "fill:"+lineColor)
	canvas.Path(pathData50, "fill:"+lineColor)
	canvas.Path(pathData51, "fill:"+lineColor)
	canvas.Path(pathData52, "fill:"+lineColor)
	canvas.Path(pathData53, "fill:"+lineColor)
	canvas.Path(pathData54, "fill:"+lineColor)
	canvas.Path(pathData55, "fill:"+lineColor)
	canvas.Path(pathData56, "fill:"+lineColor)
	canvas.Path(pathData57, "fill:"+lineColor)
	canvas.Path(pathData58, "fill:"+lineColor)
	canvas.Path(pathData59, "fill:"+lineColor)
	canvas.Path(pathData60, "fill:"+lineColor)
	canvas.Path(pathData61, "fill:"+lineColor)
	canvas.Path(pathData62, "fill:"+lineColor)
	canvas.Path(pathData63, "fill:"+lineColor)
	canvas.Path(pathData64, "fill:"+lineColor)
	canvas.Path(pathData65, "fill:"+lineColor)
	canvas.Path(pathData66, "fill:"+lineColor)
	canvas.Path(pathData67, "fill:"+lineColor)
	canvas.Path(pathData68, "fill:"+lineColor)
	canvas.Path(pathData69, "fill:"+lineColor)
	canvas.Path(pathData70, "fill:"+lineColor)
	canvas.Path(pathData71, "fill:"+lineColor)
	canvas.Path(pathData72, "fill:"+lineColor)
	canvas.Path(pathData73, "fill:"+lineColor)
	canvas.Path(pathData74, "fill:"+lineColor)
	canvas.Path(pathData75, "fill:"+lineColor)
	canvas.Path(pathData76, "fill:"+lineColor)
	canvas.Path(pathData77, "fill:"+lineColor)
	canvas.Path(pathData78, "fill:"+lineColor)
	canvas.Path(pathData79, "fill:"+lineColor)
	canvas.Path(pathData80, "fill:"+lineColor)
	canvas.Path(pathData81, "fill:"+lineColor)
	canvas.Path(pathData82, "fill:"+lineColor)
	canvas.Path(pathData83, "fill:"+lineColor)
	canvas.Path(pathData84, "fill:"+lineColor)
	canvas.Path(pathData85, "fill:"+lineColor)
	canvas.Path(pathData86, "fill:"+hexcode)
	canvas.Path(pathData87, "fill:"+hexcode)
	canvas.Path(pathData88, "fill:"+hexcode)
	canvas.Gend()
	canvas.Gend()
}
