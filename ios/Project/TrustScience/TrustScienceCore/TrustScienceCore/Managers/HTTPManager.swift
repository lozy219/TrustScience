//
//  HTTPManager.swift
//  TrustScienceCore
//
//  Created by Wang Jinghan on 10/3/19.
//  Copyright © 2019 TrustScience. All rights reserved.
//

import UIKit

public class HTTPManager: Manager {
    public func request<BuilderT: APIRequestBuildable>(builder: BuilderT, completion: @escaping (APIResult<BuilderT.ReplyT>) -> Void) {
        URLSession.shared.dataTask(
            with: builder.buildRequest(),
            completionHandler: { (data, response, error) in
                guard let data = data else {
                    completion(.error(APIError.unknown))
                    return
                }
                
                let result = builder.parseResponse(data: data)
                completion(result)
            }
        )
    }
}
