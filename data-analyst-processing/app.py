import numpy as np
import pandas as pd
from ydata_profiling import ProfileReport
from fastapi import FastAPI, HTTPException
from fastapi.responses import JSONResponse

import gc

app = FastAPI()

@app.get('/health')
async def health():
    print('health check started')
    return JSONResponse(content={
        "status": "ok"
    }, status_code=200)

@app.post('/predicts')
async def predict(data_list: List[InputData]):
    # Initialize a list to store input data for all records
    all_input_data = []

    for data in data_list:
        hour = datetime.strptime(data.Timestamp, '%Y-%m-%d %H:%M:%S').hour
        
        input_data = {
            "AS Description": data.AS_Description,
            "country": data.country,
            "state": data.state,
            "city": data.city,
            "postalcode": data.postalcode,
            "connection_type": data.connection_type,
            "coreg_path": data.coreg_path,
            "isp": data.isp,
            "Male/Female": data.male_female,
            "source": data.source,
            "subid": data.subid,
            "Age": data.Age,
            "Latitude (generated)": data.latitude,
            "Longitude (generated)": data.longitude,
            "State + City": f"{data.state} - {data.city}",
            "Source + Sub Id": f"{data.source} - {data.subid}",
            "Hour": hour,
            "Time Category": get_time_category(hour),
            "IP Address Numerized": numerize_ip(data.IP_Address)
        }

        all_input_data.append(input_data)

    df = pd.DataFrame(all_input_data)
    
    # Predict
    predictions = model.predict(df)

    predict_proba = model.predict_proba(df)
    
    # Simplify the predict_proba output
    simplified_proba = [proba[pred] for pred, proba in zip(predictions, predict_proba)]

    # Transform the predict_proba output to binary
    transformed_proba = [1 if score >= 0.5 else 0 for score in simplified_proba]

    # Prepare the response
    response = {
        # 'predictions': predictions.tolist(),
        # 'confidence_score': simplified_proba,
        'probability_score': predict_proba[:,1].tolist(),
    }

    # Return the list of predictions
    return predict_proba[:,1].tolist()