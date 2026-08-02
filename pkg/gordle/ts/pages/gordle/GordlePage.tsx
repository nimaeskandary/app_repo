import GordleView from '@app_repo/pkg_gordle/pages/gordle/GordleView'
import { useGordleViewModel } from '@app_repo/pkg_gordle/pages/gordle/useGordleViewModel'

// Connects the Gordle view to its view model.
function GordlePage() {
  const viewModel = useGordleViewModel()

  return (
    <GordleView
      cells={viewModel.cells}
      isGameComplete={viewModel.isGameComplete}
      onWordCommit={viewModel.commitWord}
    />
  )
}

export default GordlePage
