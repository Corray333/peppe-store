<template>
    <div class="product_card">
        <img :src="`http://localhost:3000/images/products/${product.img +'-'+ product.colors[0]}.jpeg`" alt="">
        <div class="card_controls">
            <div class="info_block">
                <h3>{{product.name}}</h3>
                <p>{{product.price}} r.</p>
            </div>
            <button type="button"><img src="../assets/icons/shopping-bag.svg" alt=""></button>
        </div>
    </div>
</template>


<script setup>

defineProps(['product'])


</script>

<style scoped>
.product_card{
    min-width:calc((100% - 2em) / 3);
    padding: 25px;
    border-radius: 25px;
    background-color: white;
    flex-grow: 1;
}
.product_card>img{
    width: 100%;
    border-radius: 25px;
    transition: all 0.3s;
}
.product_card:hover >img{
    filter: blur(5px);
}
.card_controls{
    display: flex;
    justify-content: space-between;
}
.info_block h3{
    text-overflow: ellipsis; 
    overflow: hidden; 
    white-space: nowrap;
    max-width: 180px;
}
.card_controls>button{
    height: fit-content;
    padding: 10px 25px;
    border: none;
    background-color: var(--accent-color);
    border-radius: 50px;
    cursor: pointer;
    transition: all 0.3s;
}
.card_controls>button:hover{
    transform: scale(1.1);
}

</style>