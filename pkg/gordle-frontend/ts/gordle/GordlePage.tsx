import GordleView from './GordleView'
import { useGordleViewModel } from './useGordleViewModel'

// Connects the Gordle view to its view model.
function GordlePage() {
  const viewModel = useGordleViewModel()

  return <GordleView cells={viewModel.cells} />
}

export default GordlePage
