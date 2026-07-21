import MuiStack, { type StackProps as MuiStackProps } from '@mui/material/Stack'

export type StackProps = MuiStackProps

export const Stack = (props: StackProps) => {
  return <MuiStack {...props} />
}
