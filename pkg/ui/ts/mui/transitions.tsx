import MuiFade, { type FadeProps as MuiFadeProps } from '@mui/material/Fade'

export type FadeProps = MuiFadeProps

export const Fade = (props: FadeProps) => {
  return <MuiFade {...props} />
}
