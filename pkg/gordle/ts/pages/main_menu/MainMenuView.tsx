import { SuccessButton } from '@app_repo/pkg_ui/mui/button'
import { Container } from '@app_repo/pkg_ui/mui/container'
import { Stack } from '@app_repo/pkg_ui/mui/stack'

export type MainMenuViewProps = {
  onPlay: () => void
}

function MainMenuView({ onPlay }: MainMenuViewProps) {
  return <Container>
    <Stack>
      <SuccessButton onClick={onPlay}>Play</SuccessButton>
    </Stack>
  </Container>
}

export default MainMenuView
