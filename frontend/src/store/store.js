import { createStore } from 'vuex'

// Create a new store instance.
const store = createStore({
  state () {
    return {
      cart: [],
    }
  },
  mutations: {
    addToCart(state, product){
        state.cart.push(product)
        localStorage.setItem('cart', JSON.stringify(state.cart))
    },
    removeFromCart(state, product){
        state.cart = state.cart.filter(item => item.id !== product.id || item.size !== product.size || item.color !== product.color)  
        localStorage.setItem('cart', JSON.stringify(state.cart))
    },
    loadCart(state){
        if(localStorage.getItem('cart')){
            state.cart = JSON.parse(localStorage.getItem('cart'))
        }
    },
    saveCart(state){
        localStorage.setItem('cart', JSON.stringify(state.cart))
    }
  }
})

export default store