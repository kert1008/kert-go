# Payment demo on AWS  
  
Pattern 1  
API Gateway --> paymentAPIHandle --> Event Bridge --> SQS --> putPaymentStepFunc --> Step Function(putPaymentStepOne --> putPaymentStepTwo) --> DynamoDB  
  
Pattern 2  
API Gateway --> paymentAPIHandle --> Event Bridge --> SQS --> paymentAdapter--> DynamoDB  
  
(paymentAPIHandle will do polling to sync payment response.)  
  
Pattern 3  
API Gateway --> setUserDynamo/getUserDynamo --> DynamoDB
