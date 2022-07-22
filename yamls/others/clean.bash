#!/bin/bash

if [[ $1 == "auto" ]]; then
    filename='autoscaler.log'
    list=(mongodb.go decider.go util.go compute_autoscaler.go cluster_feeder.go recommender.go checkpoint_writer.go metrics_client.go client.go calculator.go)
    newName=newAuto.log
elif [[ $1 == "ops" ]]; then
    filename='enterprise.log'
    list=(decider.go database.go statefulset.go mongodb_ops_request.go vertical_scaling.go restart.go parallel.go)
    newName=newOps.log
fi

# for element in "${list[@]}"; do
#     echo $element
# done

echo "" > ${newName}
n=1
while read line; do
    for element in "${list[@]}"; do
        if [[ ${line} == *"${element}"* ]]; then
            # echo $element $line
            echo "${line}" >> ${newName}
        fi
    done

    n=$((n+1))
done < $filename