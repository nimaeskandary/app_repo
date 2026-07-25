import MainMenuView from './MainMenuView'
import { useMainMenuViewModel } from './useMainMenuViewModel'

// Connects the main-menu view to its view model.
function MainMenuPage() {
  const viewModel = useMainMenuViewModel()

  return <MainMenuView onPlay={viewModel.play} />
}

export default MainMenuPage
