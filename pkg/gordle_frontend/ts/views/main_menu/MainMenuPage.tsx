import MainMenuView from '@app_repo/pkg_gordle_frontend/views/main_menu/MainMenuView'
import { useMainMenuViewModel } from '@app_repo/pkg_gordle_frontend/views/main_menu/useMainMenuViewModel'

// Connects the main-menu view to its view model.
function MainMenuPage() {
  const viewModel = useMainMenuViewModel()

  return <MainMenuView onPlay={viewModel.play} />
}

export default MainMenuPage
