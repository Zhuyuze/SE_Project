app.component('product-display2', {
    props: {
      premium: {
        type: Boolean,
        required: true
      }
    },
    template: 
    /*html*/
    `<div class="product-display">
  
  
      <div class="product-container">
        <div class="product-image">
          <img v-bind:src="image">
        </div>
        <div class="product-info">
          <h1>{{ title }}</h1>
  
          <p v-if="inStock">In Stock</p>
          <p v-else>Out of Stock</p>
  
          <p>Price: {{ price }}</p>
          <ul>
            <li v-for="detail in details">{{ detail }}</li>
          </ul>
  
          <div 
            v-for="(variant, index) in variants" 
            :key="variant.id" 
            @mouseover="updateVariant(index)" 
            class="color-circle" 
            :style="{ backgroundColor: variant.color }">{{ variant.option }}
          </div>
          
          <button 
            class="button" 
            :class="{ disabledButton: !inStock }" 
            :disabled="!inStock" 
            v-on:click="addToCart">
            Add to Cart
          </button>
        </div>
      </div>
  
  
  
    </div>`,
    data() {
      return {
          dishes:[{
            product: 'Fried Rice',
            brand: '@Mastery',
            selectedVariant: 0,
            details: ['70% rice', '15% shrimp', '10% pea & corn', '5% veg'],
            variants: [
              { id: 2234, option:'non-egg', color: 'red', image: './assets/images/item3.jpg', quantity: 50 },
              { id: 2235, option:'add-egg', color: 'grey', image: './assets/images/item3.jpg', quantity: 0 },
            ]
          },{
            product: 'Rice',
            brand: 'tery',
            selectedVariant: 0,
            details: ['70% rice', '15% shrimp', '10% pea & corn', '5% veg'],
            variants: [
              { id: 2234, option:'non-egg', color: 'red', image: './assets/images/item3.jpg', quantity: 50 },
              { id: 2235, option:'add-egg', color: 'grey', image: './assets/images/item3.jpg', quantity: 0 },
            ]
          }],
  
          product: 'Sauced Ribs',
          brand: '@Famous',
          selectedVariant: 0,
          details: ['95% pork ribs', '5% sauce'],
          variants: [
            { id: 2214, option:'non-sweet', color: 'red', image: './assets/images/item1.jpg', quantity: 50 },
            { id: 2215, option:'sweetened', color: 'red', image: './assets/images/item1.jpg', quantity: 10 },
          ],

          product: 'Peking duck',
          brand: '@Famous',
          selectedVariant: 0,
          details: ['60% roasted duck', '15% pancake', '10% sliced scallions', '10% sliced cucumbers', '5% bean paste'],
          variants: [
            { id: 2224, option:'hot-spicy', color: 'red', image: './assets/images/item2.jpg', quantity: 50 },
            { id: 2225, option:'regular', color: 'red', image: './assets/images/item2.jpg', quantity: 10 },
          ],
          
      }
    },
    methods: {
        addToCart() {
            this.$emit('add-to-cart', this.variants[this.selectedVariant].id)
        },
        updateVariant(index) {
            this.selectedVariant = index
        }
        
    },
    computed: {
        title() {
            return this.brand + ' ' + this.product
        },
        image() {
            return this.variants[this.selectedVariant].image
        },
        inStock() {
            return this.variants[this.selectedVariant].quantity
        },
        price() {
          if (this.premium) {
            //return 'Free'
            return 14.99
          }
          return 15.99
        }
    }
  })