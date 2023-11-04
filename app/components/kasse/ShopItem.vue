<script setup lang="ts">
const { storeId, itemId } = $defineProps<{
  storeId: string;
  itemId: string;
}>();

const pb = usePocketbase();

const kasseStore = useKasseStore(storeId);
let item = $ref<IShopItem>(kasseStore.getShopItem(itemId)!);

kasseStore.$subscribe((mutation, state) => {
  console.log("kasseStore.$subscribe:mutation", mutation.events);
  item = kasseStore.getShopItem(itemId)!;
  //TODO: prevent wrong events
})

function onCardClick() {
  if (item?.pfand) {
    kasseStore.removeOneCartItem(itemId);
    return;
  }
  kasseStore.addOneCardItem(itemId);
}

</script>

<template>
  <v-card v-if="item" elevation="5" :color="item.pfand ? 'yellow' : ''">
      <v-card 
        class="pa-3 d-flex justify-space-between" 
        elevation="0"
      >
        <h1>{{ item.title }}</h1>
        <v-spacer></v-spacer>
        <h1>{{ parseFloat(""+item.price).toFixed(2) }}€</h1>

      </v-card>
    <v-card 
      elevation="0"
      @click="onCardClick"
    >
      <v-img
        :src="pb.fileUrl((item as any).collectionId, item.id, item.img)"
        class="align-center justify-center w-100"
        cover
      >
        <!--<v-icon width="1000px" class="mx-auto" color="white">mdi-plus</v-icon>-->
      </v-img>
    </v-card>
      <v-card-actions class="d-flex justify-space-around">
      <v-btn 
        class="ma-0"
        icon 
        :disabled="kasseStore.getCartCount(itemId) <= 0 && !item.pfand"
        x-large 
        color="primary" 
        @click="kasseStore.removeOneCartItem(itemId)"
      >
        <v-icon x-large>mdi-minus</v-icon>
      </v-btn>
      <div class="text-center d-flex justify-center">
          {{ kasseStore.getCartCount(itemId) }}
      </div>
      <v-btn 
        class="ma-0"
        icon
        x-large 
        color="primary" 
        @click="kasseStore.addOneCardItem(itemId)"
      >
        <v-icon x-large>mdi-plus</v-icon>
      </v-btn>
    </v-card-actions>
  </v-card>
</template>

<style scoped>
h1{
  font-size: 1rem;
}
.badge {
  position: absolute;
  width: 5px;
  top: -7px;
  right: 4px;
  padding: 0px 4px;
  border-radius: 100%;
  background: #ff8e0d;
  color: white;
}
.v-img{
    height: 100px;
  }
@media (max-width: 1028px) {
  h1{
    font-size: 0.8rem;
  }
  .v-img{
    height: 50px;
  }
}
@media (max-width: 800px) {
  h1{
    font-size: 0.8rem;
  }
  .v-img{
    height: 40px;
  }
}

</style>
