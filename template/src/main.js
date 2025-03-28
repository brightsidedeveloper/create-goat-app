import './styles.css'
import './wasm_exec'

const go = new Go()
WebAssembly.instantiateStreaming(fetch('main.wasm'), go.importObject).then((result) => {
  go.run(result.instance)
})
