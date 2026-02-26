// db.books.find ( { price: { $eq: 300 } } )  

// db.books.find ( { price: { $eq: 300 } } )

// db.books.find ( { $or: [ { quantity: { $lt: 200 } }, { price: 500 } ] } )  

// db.books.find ( { $or : [ {quantity : {$gte :200}} ,{ price : {$eq : 10} } , { sellsAmt : {$lte : 50000}}]}  )

// db.students.find({ $or : [{name : {$eq : 's'}} , {address : {$nin :['asfg'] } } ]   })

// db.employees.find({name : {$in : ['charlie' , 'bob']}})

// db.employees.find({$and :[{salary : {$gte : 50000}} ,{skills : {$in : ['AWS']} } ] })

// db.orders.aggregate([
//   {
//     $lookup: {
//       from: "products",          // second collection
//       localField: "product_id",  // field in orders
//       foreignField: "_id",       // field in products
//       as: "product_details"
//     }
//   }
// ])


async function test(){
    console.log("1")
  
    await print2()

    console.log('3');
    
}

async function print2() {
    return new Promise((resolve) => {
        setTimeout(()=>{
        console.log("2");
        resolve()
        
    },1000)
    })
   
}
test()
