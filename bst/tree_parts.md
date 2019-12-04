//
//               -- 12
//              |
//      -- 10 --|
//     |        |
//     |         -- 9
// 7 --|
//     |        -- 6
//     |       |
//      -- 5 --|
//             |        -- 4
//             |       |
//              -- 3 --|
//                     |
//                      -- 1
//
//
//
//
//
//



                          ┌ ─ 14  ─ ┤

                          │

                ┌ ─ 12  ─ ┤

                │

     ┌ ─  10  ─ ┤

     │          │

     │          └ ─ 9 ─ ┤

     │                  │

     │                  └ ─ 8 ─ ┤

     │

 7 ─ ┤

     
     │         ┌ ─ 6 ─ ┤

     │         │

     └ ─  5  ─ ┤

               │

               │       ┌ ─ 4 ─ ┤

               │       │

               └ ─ 3 ─ ┤

                       │

                       └ ─ 1 ─ ┤


             ┌─14─┤
             │
        ┌─12─┘
        │
   ┌─10─┤
   │    │
   │    └─9 ─┐
   │         │
   │         └─8 ─┤
   │
 7─┤
   │
   │    ┌─6 ─┤
   │    │
   └─5 ─┤
        │
        │    ┌─4 ─┤
        │    │
        └─3 ─┤
             │
             └─1 ─┤



                        ┌─ <14> ─┤
                        │
               ┌─ <12> ─┘
               │
      ┌─ <10> ─┤
      │        │
      │        └─ <9 > ─┐
      │                 │
      │                 └─ <8 > ─┤
      │
<7 > ─┤
      │
      │        ┌─ <6 > ─┤
      │        │     
      └─ <5 > ─┤     
               │     
               │        ┌─ <4 > ─┤
               │        │
               └─ <3 > ─┤
                        │
                        └─ <1 > ─┤



                   ┌─<14>
                   │
            ┌─<12>─┘
            │
     ┌─<10>─┤
     │      │
     │      └─<9 >─┐
     │             │
     │             └─<8 >
     │
<7 >─┤
     │
     │      ┌─<6 >
     │      │    
     └─<5 >─┤    
            │    
            │      ┌─<4 >
            │      │
            └─<3 >─┤
                   │
                   └─<1 >


                   ┌─[14]
                   │
            ┌─[12]─┘
            │
     ┌─[10]─┤
     │      │
     │      └─[ 9]─┐
     │             │
     │             └─[ 8]
     │
[ 7]─┤
     │
     │      ┌─[ 6]
     │      │    
     └─[ 5]─┤    
            │    
            │      ┌─[ 4]
            │      │
            └─[ 3]─┤
                   │
                   └─[ 1]






                   ┌─[18]
                   │
            ┌─[16]─┘
            │
     ┌─[14]─┤
     │      │
     │      │      ┌─[13]
     │      │      │
     │      └─[12]─┤
     │             │
     │             │      ┌─[11]
     │             │      │
     │             └─[10]─┤
     │                    │
     │                    └─[ 9]
     │
[ 7]─┤
     │
     │      ┌─[ 6]
     │      │    
     └─[ 5]─┤    
            │    
            │      ┌─[ 4]
            │      │
            └─[ 3]─┤
                   │
                   └─[ 1]







                   ┌─[18]
            ┌─[16]─┘
     ┌─[14]─┤
     │      │      ┌─[13]
     │      └─[12]─┤
     │             │      ┌─[11]
     │             └─[10]─┤
     │                    └─[ 9]
[ 7]─┤
     │
     │
     │
     │
     │      ┌─[ 6]
     └─[ 5]─┤    
            │      ┌─[ 4]
            └─[ 3]─┤
                   └─[ 1]


### Ready prototype
                   ┌─[18]
            ┌─[16]─┘
     ┌─[14]─┤
     │      │      ┌─[13]
     │      └─[12]─┤
     │             │      ┌─[11]
     │             └─[10]─┤
     │                    └─[ 9]
[ 7]─┤
     │
     │
     │
     │
     │      ┌─[ 6]
     └─[ 5]─┤    
            │      ┌─[ 4]
            └─[ 3]─┤
                   └─[ 1]


### Ready prototype with blanks
                         ┌ ─ [18]

              ┌ ─ [16] ─ ┘

     ┌─[14] ─ ┤ 

     │        │          ┌ ─ [13]

     │        └ ─ [12] ─ ┤ 

     │                   │          ┌ ─ [11]

     │                   └ ─ [10] ─ ┤ 

     │                              └ ─ [ 9]

[ 7]─┤

     │

     │      ┌─[ 6]

     └─[ 5]─┤    

            │      ┌─[ 4]

            └─[ 3]─┤

                   └─[ 1]



│ 2502 179 BOX DRAWINGS LIGHT VERTICAL

┤ 2524 180 BOX DRAWINGS LIGHT VERTICAL AND LEFT

┌ 250C 218 BOX DRAWINGS LIGHT DOWN AND RIGHT

┘ 2518 217 BOX DRAWINGS LIGHT UP AND LEFT

┐ 2510 191 BOX DRAWINGS LIGHT DOWN AND LEFT

└ 2514 192 BOX DRAWINGS LIGHT UP AND RIGHT

─ 2500 196 BOX DRAWINGS LIGHT HORIZONTAL


VERT       = "│" //2502 179 
VERT_LEFT  = "┤" //2524 180 
DOWN_RIGHT = "┌" //250C 218 
UP_LEFT    = "┘" //2518 217 
DOWN_LEFT  = "┐" //2510 191 
UP_RIGHT   = "└" //2514 192 
HORIZONTAL = "─" 2500 196 



###############
⌐ 2310 169 

¬ 00AC 170 

┴ 2534 193

┬ 252C 194

├ 251C 195

─ 2500 196

┼ 253C 197



// 1 0
// 2 0   0*5
// 3 5   1*5
// 4 10  2*5
// 5 15  3*5
// 6 20  4*5
// 7 25  5*5


0123012301230123
1
└─── 1
     ├─── 1            5
     ├─── 1            5
     ├─── 1            5
     ├─── 1            5
     ├─── 1            5
     └─── 1            5
          └─── 1       10
          └─── 1       10
               └─── 1  15
          └─── 1
     └─── 1









1           
└─── 3                                   node                
     ├─── 5                              true node                       
     └─── 8                              true node           
└─── 6                                   node           
     ├─── 17                             true node            
     ├─── 18                             true node            
     └─── 19                             true node            
          └─── 25                        true true node           
               └─── 26                   true true fals node                
                    └─── 306             ture true                            
                         ├─── 317                                                  
                         ├─── 318                                                  
                         └─── 319                                                  
                              ├─── 325                                                  
                              └─── 329                                                  
                         ├─── 339                                                  
                         └─── 349                                                  
          └─── 29                                                  
     ├─── 39                                                  
     └─── 49                                                  
└─── 73                                                  
     ├─── 85                                                  
     └─── 88                                                  
          └─── 91                                                  
               └─── 94                                                  
└─── 7                                                  









1                                                                      last: false
└─── 3                                                                 last: false
     ├─── 5                                                            last: false
     └─── 8                                                            last: true
└─── 6                                                                 last: false
     ├─── 17                                                           last: false
     ├─── 18                                                           last: false
     └─── 19                                                           last: false
          └─── 25                                                      last: false
               └─── 26                                                 last: true
                    └─── 306                                           last: true
                         ├─── 317                                      last: false
                         ├─── 318                                      last: false
                         └─── 319                                      last: false
                              ├─── 325                                 last: false
                              └─── 329                                 last: true
                         ├─── 339                                      last: false
                         └─── 349                                      last: true
          └─── 29                                                      last: true
     ├─── 39                                                           last: false
     └─── 49                                                           last: true
└─── 73                                                                last: false
     ├─── 85                                                           last: false
     └─── 88                                                           last: true
          └─── 91                                                      last: true
               └─── 94                                                 last: true
└─── 7                                                                 last: true



                                            should print predecesor │?
1                                                                              last: false
└─── 3                                      node                               last: false
│    ├─── 5                                 true node                          last: false
│    └─── 8                                 true node                          last: true
└─── 6                                      node                               last: false
│    ├─── 17                                true node                          last: false
│    ├─── 18                                true node                          last: false
│    └─── 19                                true node                          last: false
│    │    └─── 25                           true true node                     last: false
│    │    │    └─── 26                      true true true node                last: true
│    │    │    *    └─── 306                true true true fals node           last: true
│    │    │    *    *    ├─── 317           true true true fals fals node      last: false
│    │    │    *    *    ├─── 318           true true true fals fals node      last: false
│    │    │    *    *    └─── 319           true true true fals fals node      last: false
│    │    │    *    *    │    ├─── 325      true true true fals fals true node last: false
│    │    │    *    *    │    └─── 329      true true true fals fals true node last: true
│    │    │    *    *    ├─── 339           true true true fals fals node      last: false
│    │    │    *    *    └─── 349           true true true fals fals node      last: true
│    │    └─── 29                           true true node                     last: true
│    ├─── 39                                true node                          last: false
│    └─── 49                                true node                          last: true
└─── 73                                     node                               last: false
│    ├─── 85                                true node                          last: false
│    └─── 88                                true node                          last: true
│    *    └─── 91                           true fals node                     last: true
│    *    *    └─── 94                      true fals fals node                last: true
└─── 7                                      node                               last: true




                                            should print predecesor │?
1                                      r 
└─── 3                                 r [ 3]fals 
│    ├─── 5                            r [ 3]fals [ 5]fals
│    └─── 8                            r [ 3]fals [ 8]true
└─── 6                                 r [ 6]fals
│    ├─── 17                           r [ 6]fals [17]fals
│    ├─── 18                           r [ 6]fals [18]fals
│    └─── 19                           r [ 6]fals [19]fals
│    │    └─── 25                      r [ 6]fals [19]fals [25]fals
│    │    │    └─── 26                 r [ 6]fals [19]fals [25]fals [26]true
│    │    │    *    └─── 306           r [ 6]fals [19]fals [25]fals [26]true [306]true
│    │    │    *    *    ├─── 317      r [ 6]fals [19]fals [25]fals [26]true [306]true [317]fals
│    │    │    *    *    ├─── 318      r [ 6]fals [19]fals [25]fals [26]true [306]true [318]fals
│    │    │    *    *    └─── 319      r [ 6]fals [19]fals [25]fals [26]true [306]true [319]fals
│    │    │    *    *    │    ├─── 325 r [ 6]fals [19]fals [25]fals [26]true [306]true [319]fals [325]false
│    │    │    *    *    │    └─── 329 r [ 6]fals [19]fals [25]fals [26]true [306]true [319]fals [329]true
│    │    │    *    *    ├─── 339      r [ 6]fals [19]fals [25]fals [26]true [306]true [339]fals
│    │    │    *    *    └─── 349      r [ 6]fals [19]fals [25]fals [26]true [306]true [349]true
│    │    └─── 29                      r [ 6]fals [19]fals [29]true
│    ├─── 39                           r [ 6]fals [39]fals
│    └─── 49                           r [ 6]fals [49]true
└─── 73                                r [73]fals
│    ├─── 85                           r [73]fals [85]fals
│    └─── 88                           r [73]fals [88]true
│    *    └─── 91                      r [73]fals [88]true [91]true
│    *    *    └─── 94                 r [73]fals [88]true [91]true [94]true
└─── 7                                 r [ 7]true 



[] [1]
[false] [3]
[false false] [5]
[false true] [8]
[false] [6]
[false false] [17]
[false false] [18]
[false false] [19]
[false false false] [25]
[false false false true] [26]
[false false false true true] [306]
[false false false true true false] [317]
[false false false true true false] [318]
[false false false true true false] [319]
[false false false true true false false] [325]
[false false false true true false true] [329]
[false false false true true false] [339]
[false false false true true true] [349]
[false false true] [29]
[false false] [39]
[false true] [49]
[false] [73]
[false false] [85]
[false true] [88]
[false true true] [91]
[false true true true] [94]
[true] [7]


last element removed
[] [1]
[] [3]
[false] [5]
[false] [8]
[false] [6]
[false] [17]
[false] [18]
[false] [19]
[false false] [25]
[false false false] [26]
[false false false true] [306]
[false false false true true] [317]
[false false false true true] [318]
[false false false true true] [319]
[false false false true true false] [325]
[false false false true true false] [329]
[false false false true true] [339]
[false false false true true] [349]
[false false] [29]
[false] [39]
[false] [49]
[] [73]
[false] [85]
[false] [88]
[false true] [91]
[false true true] [94]
[] [7]


func getter(laster []bool) {
last, rest = laster[len(laster)-1], laster[:len(laster)-1]



}





