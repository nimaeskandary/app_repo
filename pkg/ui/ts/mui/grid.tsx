import MuiGrid, { type GridProps as MuiGridProps } from '@mui/material/Grid'

export type GridProps = MuiGridProps

export const Grid = (props: GridProps) => {
  return <MuiGrid {...props} />
}
