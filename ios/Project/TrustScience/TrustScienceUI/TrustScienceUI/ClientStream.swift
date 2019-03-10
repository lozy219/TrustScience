//
//  ClientStream.swift
//  TrustScience
//
//  Created by Wang Jinghan on 10/3/19.
//  Copyright © 2019 TrustScience. All rights reserved.
//

import UIKit

public class ClientStream {
    public struct Dependencies {
        public let window: UIWindow
        public init(window: UIWindow) {
            self.window = window
        }
    }
    
    public static private(set) var shared: ClientStream!
    static var session: ClientSession { return shared.session }
    public static func setup(with dependencies: Dependencies) {
        guard shared == nil else {
            assertionFailure()
            return
        }
        
        shared = ClientStream(dependencies: dependencies)
    }
    
    private let dependencies: Dependencies
    private(set) var session: ClientSession
    
    public init(dependencies: Dependencies) {
        self.dependencies = dependencies
        self.session = ClientSession(window: dependencies.window)
        self.session.delegate = self
    }
    
    public func startSession() {
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
