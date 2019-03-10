//
//  ClientStream.swift
//  TrustScience
//
//  Created by Wang Jinghan on 10/3/19.
//  Copyright © 2019 TrustScience. All rights reserved.
//

import UIKit

class ClientStream {
    struct Dependencies {
        let window: UIWindow
    }
    
    static private(set) var shared: ClientStream!
    static var session: ClientSession { return shared.session }
    static func setup(with dependencies: Dependencies) {
        guard shared == nil else {
            assertionFailure()
            return
        }
        
        shared = ClientStream(dependencies: dependencies)
    }
    
    private let dependencies: Dependencies
    private(set) var session: ClientSession
    
    init(dependencies: Dependencies) {
        self.dependencies = dependencies
        self.session = ClientSession(window: dependencies.window)
        self.session.delegate = self
    }
    
    func startSession() {
        self.session.start()
    }
}

extension ClientStream: ClientSessionDelegate {
    func clientSessionDidEnd(_ session: ClientSession) {
        self.session = ClientSession(window: dependencies.window)
        self.session.delegate = self
        
        self.startSession()
    }
}
