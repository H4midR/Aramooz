<template>
    <div>
        <v-list-tile v-for="(choice,n) in (choicesList.choices)" :key="'q_'+n+'_ch'">
              <v-list-tile-avatar>
                <v-btn small depressed color="success" fab>
                  {{n+1}}
                </v-btn>
              </v-list-tile-avatar>
              <v-list-tile-content>
                <v-list-tile-title :class=" (choicesList.dir=='rtl') ? 'text-xs-right': 'text-xs-left'" v-html="choice.title"></v-list-tile-title>
              </v-list-tile-content>
        </v-list-tile>
    </div>
</template>

<script>
export default {
porps:{
    choicesList:Array
}
}
</script>

<style>

</style>
