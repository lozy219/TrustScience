//
//  APIRequestBuildable.swift
//  TrustScienceCore
//
//  Created by Wang Jinghan on 10/3/19.
//  Copyright © 2019 TrustScience. All rights reserved.
//

import Foundation

public protocol APIRequestBuildable {
    associatedtype ReplyT
    
    func buildRequest() -> URLRequest
    func parseResponse(data: Data) -> APIResult<ReplyT>
}

public enum APIResult<ReplyT> {
    case success(ReplyT)
    case error(Error)
}

public enum APIError: Error {
    case unknown
}
