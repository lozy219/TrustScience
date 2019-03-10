//
//  ResultRequestBuilder.swift
//  TrustScienceCore
//
//  Created by Wang Jinghan on 10/3/19.
//  Copyright © 2019 TrustScience. All rights reserved.
//

import UIKit

struct ResultRequestBuilder: APIRequestBuildable {
    typealias ReplyT = Reply
    
    private let url = URL(string: "http://uygnim.com:8734/result")!
    
    func buildRequest() -> URLRequest {
        var request = URLRequest(url: url)
        request.httpMethod = "GET"
        return request
    }
    
    func parseResponse(data: Data) -> APIResult<Reply> {
        do {
            let reply = try JSONDecoder().decode(Reply.self, from: data)
            return .success(reply)
        } catch {
            return .error(error)
        }
    }
}

extension ResultRequestBuilder {
    struct Reply: Codable {
        let current: String
        let preview: String
        let result: [UInt32]
    }
}
