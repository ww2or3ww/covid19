export const single: string = '#3088BF'
export const double: string[] = ['#006400', '#1B75BC']
export const triple: string[] = ['#00441B', '#006400', '#1B75BC']
export const quadruple: string[] = ['#00441B', '#006400', '#1B75BC', '#505B00']

export type SurfaceStyle = {
  strokeColor: string
  fillColor: string
}

const surfaceStyleA: SurfaceStyle = {
  strokeColor: '#3088BF',
  fillColor: '#3088BF'
}

const surfaceStyleB: SurfaceStyle = {
  strokeColor: '#3088BF',
  fillColor: '#3088BF'
}

const surfaceStyleC: SurfaceStyle = {
  strokeColor: '#3088BF',
  fillColor: '#3088BF'
}

export function getGraphSeriesStyle(seriesLength: number) {
  switch (seriesLength) {
    case 1:
      return [surfaceStyleB]
    case 2:
      return [surfaceStyleA, surfaceStyleC]
    default:
      return [surfaceStyleA, surfaceStyleB, surfaceStyleC]
  }
}
