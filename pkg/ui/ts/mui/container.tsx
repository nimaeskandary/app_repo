import MuiContainer, { type ContainerProps as MuiContainerProps } from '@mui/material/Container'

export type ContainerProps = MuiContainerProps

export const Container = (props: ContainerProps) => {
  return <MuiContainer {...props} />
}
