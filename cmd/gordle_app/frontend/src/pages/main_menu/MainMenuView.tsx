import { SuccessButton } from "@/pkg/ui/ts/mui/button";
import { Container } from "@/pkg/ui/ts/mui/container";
import { Stack } from "@/pkg/ui/ts/mui/stack";
import { Events } from '@wailsio/runtime';

function MainMenuView() {
    return <Container>
        <Stack>
            <SuccessButton onClick={() => Events.Emit('gordle_app:navigate:path', '/gordle')}>Play</SuccessButton>
        </Stack>
    </Container>
}

export default MainMenuView
